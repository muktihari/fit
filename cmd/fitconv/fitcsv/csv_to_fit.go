// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fitcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type CSVToFITConv struct {
	enc       *encoder.Encoder
	streamEnc *encoder.StreamEncoder
	csv       *csv.Reader

	fit proto.FIT

	fieldsArray          [255]proto.Field
	developerFieldsArray [255]proto.DeveloperField
	protoValuesArray     [255]proto.Value

	fieldDescriptions []*mesgdef.FieldDescription

	unknownMesg         int
	unknownField        int
	unknwonDynamicField int
	seq                 int
	line                int
	col                 int
}

// NewCSVToFITConv creates a new CSV to FIT converter.
func NewCSVToFITConv(fitWriter io.Writer, csvReader io.Reader) *CSVToFITConv {
	enc := encoder.New(fitWriter,
		encoder.WithProtocolVersion(proto.V2),
		encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15),
		encoder.WithMessageValidator(
			encoder.NewMessageValidator(
				encoder.ValidatorWithPreserveInvalidValues(),
			),
		),
	)

	csvr := csv.NewReader(csvReader)
	csvr.FieldsPerRecord = -1 // dynamic number of fields
	streamEnc, _ := enc.StreamEncoder()
	return &CSVToFITConv{
		enc:       enc,
		streamEnc: streamEnc,
		csv:       csvr,
	}
}

type CSVToFITConvInfo struct {
	UnknownMesg         int
	UnknownField        int
	UnknownDynamicField int
	Sequence            int
}

func (c *CSVToFITConv) ResultInfo() CSVToFITConvInfo {
	return CSVToFITConvInfo{
		UnknownMesg:         c.unknownMesg,
		UnknownField:        c.unknownField,
		UnknownDynamicField: c.unknwonDynamicField,
		Sequence:            c.seq,
	}
}

func (c *CSVToFITConv) Convert() error {
	if err := c.convert(); err != nil {
		return fmt.Errorf("[line: %d, col: %d]: %w", c.line, c.col, err)
	}
	return nil
}

func (c *CSVToFITConv) convert() error {
	for {
		c.line++

		record, err := c.csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("csv.Read: %w", err)
		}

		if len(record) < 3 {
			continue
		}

		var (
			header   = record[0]
			_        = record[1] // local message num
			mesgName = record[2]
		)

		switch header {
		case "Type", "Definition":
			continue
		case "Data":
			mesgNum, ok := mesgNumLookup[mesgName]
			if !ok {
				digits := strings.Map(func(r rune) rune {
					if !unicode.IsDigit(r) {
						return -1
					}
					return r
				}, mesgName)
				if digits == "" {
					c.unknownMesg++
					continue
				}
				// Unknown has number, try parsing it as MesgNum.
				v, err := strconv.ParseUint(digits, 10, 16)
				if err != nil {
					return fmt.Errorf("could not parse unknown mesgNum %q: %w", mesgName, err)
				}
				mesgNum = typedef.MesgNum(v)
			}

			if mesgNum == mesgnum.FileId {
				if c.seq != 0 {
					if err := c.finalize(); err != nil {
						return fmt.Errorf("could not finalize: %w", err)
					}
				}
				c.seq++
			}

			mesg, err := c.createMesg(mesgNum, record)
			if err != nil {
				return fmt.Errorf("could not create mesg: num: %q (%d): %w", mesgNum, mesgNum, err)
			}

			if mesg.Num == mesgnum.FieldDescription {
				c.fieldDescriptions = append(c.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
			}

			if c.streamEnc == nil {
				mesg.Fields = append(mesg.Fields[:0:0], mesg.Fields...)
				mesg.DeveloperFields = append(mesg.DeveloperFields[:0:0], mesg.DeveloperFields...)
				c.fit.Messages = append(c.fit.Messages, mesg)
				continue
			}

			if err = c.streamEnc.WriteMessage(&mesg); err != nil {
				return err
			}
		}
	}

	return c.finalize()
}

func (c *CSVToFITConv) finalize() error {
	if c.streamEnc == nil {
		if err := c.enc.Encode(&c.fit); err != nil {
			return err
		}
		c.fit = proto.FIT{} // reset
		return nil
	}
	return c.streamEnc.SequenceCompleted()
}

type dynamicFieldRef struct {
	index int
	name  string
	value string
}

var placeholderField = proto.Field{FieldBase: &proto.FieldBase{Num: 255, Name: "fitconv_field_placeholder"}}

func (c *CSVToFITConv) createMesg(num typedef.MesgNum, record []string) (proto.Message, error) {
	c.col = 0

	mesg := proto.Message{
		Num:             num,
		Fields:          c.fieldsArray[:0],
		DeveloperFields: c.developerFieldsArray[:0],
	}

	if len(record) < 6 {
		return mesg, nil
	}

	var dynamicFieldRefs []dynamicFieldRef

	for i := 3; i+2 < len(record); i += 3 {
		c.col = i
		var (
			fieldName = record[i+0]
			strValue  = record[i+1]
			units     = record[i+2]
		)
		if fieldName == "" {
			continue
		}

		var recoverableUnknownField bool
		var fieldNum byte
		var ok bool
		if int(num) < len(fieldNumLookup) {
			fieldNum, ok = fieldNumLookup[num][fieldName]
		}
		if !ok {
			if strings.HasPrefix(fieldName, factory.NameUnknown) {
				digits := strings.Map(func(r rune) rune {
					if !unicode.IsDigit(r) {
						return -1
					}
					return r
				}, fieldName)
				if digits == "" {
					c.unknownField++
					continue
				}
				v, err := strconv.ParseUint(digits, 10, 8)
				if err != nil {
					return mesg, fmt.Errorf("could not parse unknown fieldNum %q: %w", fieldName, err)
				}
				fieldNum = byte(v)
				recoverableUnknownField = true
				ok = true
			}
		}

		if ok {
			field, err := c.createField(num, fieldNum, strValue, units, recoverableUnknownField)
			if err != nil {
				return mesg, fmt.Errorf("could not create field: num: %q (%d): %w", fieldName, fieldNum, err)
			}
			mesg.Fields = append(mesg.Fields, field)
			continue
		}

		devField, err := c.createDeveloperField(fieldName, strValue, units)
		if err != nil {
			return mesg, fmt.Errorf("could not create developer field: %w", err)
		}
		if devField.Value.Type() != proto.TypeInvalid {
			mesg.DeveloperFields = append(mesg.DeveloperFields, devField)
			return mesg, nil
		}

		// If the field cannot be found in fieldNumLookup and isn't a valid developer field,
		// let's marked as a dynamic field and revert it later.
		dynamicFieldRefs = append(dynamicFieldRefs, dynamicFieldRef{
			index: len(mesg.Fields),
			name:  fieldName,
			value: strValue,
		})
		mesg.Fields = append(mesg.Fields, placeholderField)
	}

	for _, ref := range dynamicFieldRefs {
		if err := c.revertSubFieldSubtitution(&mesg, ref); err != nil {
			return mesg, err
		}
	}

	// Remove remaining field placeholders if it can't be subtituted.
	var valid int
	for i := range mesg.Fields {
		if mesg.Fields[i].Name == placeholderField.Name {
			continue
		}
		if i != valid {
			mesg.Fields[i], mesg.Fields[valid] = mesg.Fields[valid], mesg.Fields[i]
		}
		valid++
	}
	mesg.Fields = mesg.Fields[:valid]

	removeExpandedComponents(&mesg)

	return mesg, nil
}

func (c *CSVToFITConv) createField(mesgNum typedef.MesgNum, num byte, strValue, units string, recoverableUnknownField bool) (field proto.Field, err error) {
	field = factory.CreateField(mesgNum, num)
	if recoverableUnknownField {
		field.BaseType = basetype.FromString(units)
		field.Type = profile.ProfileTypeFromBaseType(field.BaseType)
	}

	sliceValues := strings.Split(strValue, "|")
	if len(sliceValues) != 1 || field.Array {
		protoValues := c.protoValuesArray[:0]
		for i := range sliceValues {
			value, err := parseValue(
				sliceValues[i],
				field.BaseType,
				field.Type,
				field.Scale,
				field.Offset,
				units,
			)
			if err != nil {
				return field, fmt.Errorf("[%d]: could not parse %q into %s: %v",
					i, sliceValues[i], field.BaseType.GoType(), err)
			}
			protoValues = append(protoValues, value)
		}
		field.Value = packValues(protoValues)
		return
	}

	field.Value, err = parseValue(
		strValue,
		field.BaseType,
		field.Type,
		field.Scale,
		field.Offset,
		units,
	)
	if err != nil {
		err = fmt.Errorf("could not parse %q into %s: %v",
			strValue, field.BaseType.GoType(), err)
	}

	return
}

func (c *CSVToFITConv) createDeveloperField(name, strValue, units string) (devField proto.DeveloperField, err error) {
	var fieldDesc *mesgdef.FieldDescription
	for i := range c.fieldDescriptions {
		if strings.Join(c.fieldDescriptions[i].FieldName, "|") == name {
			fieldDesc = c.fieldDescriptions[i]
			break
		}
	}
	if fieldDesc == nil {
		return
	}

	scale, offset := 1.0, 0.0 // default
	if fieldDesc.Scale != basetype.Uint8Invalid {
		scale = float64(fieldDesc.Scale)
	}
	if fieldDesc.Offset != basetype.Sint8Invalid {
		offset = float64(fieldDesc.Offset)
	}

	devField = proto.DeveloperField{
		DeveloperDataIndex: fieldDesc.DeveloperDataIndex,
		Num:                fieldDesc.FieldDefinitionNumber,
	}

	sliceValues := strings.Split(strValue, "|")
	if len(sliceValues) != 1 {
		protoValues := c.protoValuesArray[:0]
		for i := range sliceValues {
			value, err := parseValue(
				sliceValues[i],
				fieldDesc.FitBaseTypeId,
				profile.ProfileTypeFromBaseType(fieldDesc.FitBaseTypeId),
				scale,
				offset,
				units,
			)
			if err != nil {
				return devField, fmt.Errorf("%q: [%d]: could not parse %q into %s ",
					name, i, sliceValues[i], fieldDesc.FitBaseTypeId.GoType())
			}
			protoValues = append(protoValues, value)
		}
		devField.Value = packValues(protoValues)
		return
	}

	devField.Value, err = parseValue(
		strValue,
		fieldDesc.FitBaseTypeId,
		profile.ProfileTypeFromBaseType(fieldDesc.FitBaseTypeId),
		scale,
		offset,
		units,
	)
	if err != nil {
		err = fmt.Errorf("%q: could not parse %q into %s",
			name, strValue, fieldDesc.FitBaseTypeId.GoType())
	}

	return
}

func (c *CSVToFITConv) revertSubFieldSubtitution(mesgRef *proto.Message, ref dynamicFieldRef) (err error) {
	mesgLookup := factory.CreateMesg(mesgRef.Num)
	for _, fieldRef := range mesgLookup.Fields {
		for _, subField := range fieldRef.SubFields {
			if subField.Name != ref.name {
				continue
			}
			for _, smap := range subField.Maps {
				valRef, ok := convertToInt64(mesgRef.FieldValueByNum(smap.RefFieldNum))
				if !ok {
					continue
				}

				if smap.RefFieldValue == valRef {
					fieldRef.Value, err = parseValue(
						ref.value,
						fieldRef.BaseType,
						fieldRef.Type,
						fieldRef.Scale,
						fieldRef.Offset,
						fieldRef.Units,
					)
					if err != nil {
						return err
					}

					mesgRef.Fields[ref.index] = fieldRef // replace field placeholder
					return nil
				}
			}
		}
	}

	c.unknwonDynamicField++
	return nil
}

func removeExpandedComponents(mesg *proto.Message) {
	fields := map[byte]struct{}{}
	for _, field := range mesg.Fields {
		fields[field.Num] = struct{}{}
	}

	// NOTE: brute-force checking all possible candidates, if too slow optimize this.
	removeCandidates := map[byte]struct{}{}
	for _, field := range mesg.Fields {
		for _, subField := range field.SubFields {
			for _, component := range subField.Components {
				_, ok := fields[component.FieldNum]
				if ok {
					removeCandidates[component.FieldNum] = struct{}{}
				}
			}
		}
		for _, component := range field.Components {
			_, ok := fields[component.FieldNum]
			if ok {
				removeCandidates[component.FieldNum] = struct{}{}
			}
		}
	}

	for num := range removeCandidates {
		mesg.RemoveFieldByNum(num)
	}
}

func parseValue(strValue string, baseType basetype.BaseType, profileType profile.ProfileType, scale, offset float64, units string) (value proto.Value, err error) {
	if units == "degrees" && baseType == basetype.Sint32 { // Special case
		degrees, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return value, err
		}
		return proto.Int32(semicircles.ToSemicircles(degrees)), nil
	}

	if profileType == profile.Bool {
		v, err := strconv.ParseUint(strValue, 0, 8)
		if err != nil {
			return value, err
		}
		return proto.Bool(typedef.Bool(v)), nil
	}

	var scaledValue float64
	var isScaled bool = baseType != basetype.String && strings.Contains(strValue, ".")
	if isScaled {
		scaledValue, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return value, fmt.Errorf("try")
		}
		scaledValue = scaleoffset.Discard(scaledValue, scale, offset)
	}

	switch baseType {
	case basetype.Enum, basetype.Byte,
		basetype.Uint8, basetype.Uint8z:
		if isScaled {
			return proto.Uint8(uint8(scaledValue)), nil
		}
		var v uint64
		v, err = strconv.ParseUint(strValue, 0, 8)
		if err != nil {
			return
		}
		return proto.Uint8(uint8(v)), nil
	case basetype.Sint8:
		if isScaled {
			return proto.Int8(int8(scaledValue)), nil
		}
		var v int64
		v, err = strconv.ParseInt(strValue, 0, 8)
		if err != nil {
			return
		}
		return proto.Int8(int8(v)), nil
	case basetype.Sint16:
		if isScaled {
			return proto.Int16(int16(scaledValue)), nil
		}
		var v int64
		v, err = strconv.ParseInt(strValue, 0, 16)
		if err != nil {
			return
		}
		return proto.Int16(int16(v)), nil
	case basetype.Uint16, basetype.Uint16z:
		if isScaled {
			return proto.Uint16(uint16(scaledValue)), nil
		}
		var v uint64
		v, err = strconv.ParseUint(strValue, 0, 16)
		if err != nil {
			return
		}
		return proto.Uint16(uint16(v)), nil
	case basetype.Sint32:
		if isScaled {
			return proto.Int32(int32(scaledValue)), nil
		}
		var v int64
		v, err = strconv.ParseInt(strValue, 0, 32)
		if err != nil {
			return
		}
		return proto.Int32(int32(v)), nil
	case basetype.Uint32, basetype.Uint32z:
		if isScaled {
			return proto.Uint32(uint32(scaledValue)), nil
		}
		var v uint64
		v, err = strconv.ParseUint(strValue, 0, 32)
		if err != nil {
			return
		}
		return proto.Uint32(uint32(v)), nil
	case basetype.String:
		return proto.String(strValue), nil
	case basetype.Float32:
		if isScaled {
			return proto.Float32(float32(scaledValue)), nil
		}
		var v float64
		v, err = strconv.ParseFloat(strValue, 32)
		if err != nil {
			return
		}
		return proto.Float32(float32(v)), nil
	case basetype.Float64:
		if isScaled {
			return proto.Float64(scaledValue), nil
		}
		var v float64
		v, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return
		}
		return proto.Float64(v), nil
	case basetype.Sint64:
		if isScaled {
			return proto.Int64(int64(scaledValue)), nil
		}
		var v int64
		v, err = strconv.ParseInt(strValue, 0, 64)
		if err != nil {
			return
		}
		return proto.Int64(v), nil
	case basetype.Uint64, basetype.Uint64z:
		if isScaled {
			return proto.Uint64(uint64(scaledValue)), nil
		}
		var v uint64
		v, err = strconv.ParseUint(strValue, 0, 64)
		if err != nil {
			return
		}
		return proto.Uint64(v), nil
	}

	return
}

func packValues(vals []proto.Value) proto.Value {
	if len(vals) == 0 {
		return proto.Value{}
	}
	switch vals[0].Type() {
	case proto.TypeBool:
		values := make([]typedef.Bool, len(vals))
		for i := range vals {
			values[i] = vals[i].Bool()
		}
		return proto.SliceBool(values)
	case proto.TypeInt8:
		values := make([]int8, len(vals))
		for i := range vals {
			values[i] = vals[i].Int8()
		}
		return proto.SliceInt8(values)
	case proto.TypeUint8:
		values := make([]uint8, len(vals))
		for i := range vals {
			values[i] = vals[i].Uint8()
		}
		return proto.SliceUint8(values)
	case proto.TypeInt16:
		values := make([]int16, len(vals))
		for i := range vals {
			values[i] = vals[i].Int16()
		}
		return proto.SliceInt16(values)
	case proto.TypeUint16:
		values := make([]uint16, len(vals))
		for i := range vals {
			values[i] = vals[i].Uint16()
		}
		return proto.SliceUint16(values)
	case proto.TypeInt32:
		values := make([]int32, len(vals))
		for i := range vals {
			values[i] = vals[i].Int32()
		}
		return proto.SliceInt32(values)
	case proto.TypeUint32:
		values := make([]uint32, len(vals))
		for i := range vals {
			values[i] = vals[i].Uint32()
		}
		return proto.SliceUint32(values)
	case proto.TypeInt64:
		values := make([]int64, len(vals))
		for i := range vals {
			values[i] = vals[i].Int64()
		}
		return proto.SliceInt64(values)
	case proto.TypeUint64:
		values := make([]uint64, len(vals))
		for i := range vals {
			values[i] = vals[i].Uint64()
		}
		return proto.SliceUint64(values)
	case proto.TypeFloat32:
		values := make([]float32, len(vals))
		for i := range vals {
			values[i] = vals[i].Float32()
		}
		return proto.SliceFloat32(values)
	case proto.TypeFloat64:
		values := make([]float64, len(vals))
		for i := range vals {
			values[i] = vals[i].Float64()
		}
		return proto.SliceFloat64(values)
	case proto.TypeString:
		values := make([]string, len(vals))
		for i := range vals {
			values[i] = vals[i].String()
		}
		return proto.SliceString(values)
	}
	return proto.Value{}
}

// convertToInt64 converts any integer value of val to int64, if val is non-integer value return false.
// This should be keep updated to align with proto.convertToInt64.
func convertToInt64(val proto.Value) (int64, bool) {
	switch val.Type() {
	case proto.TypeInt8:
		return int64(val.Int8()), true
	case proto.TypeUint8:
		return int64(val.Uint8()), true
	case proto.TypeInt16:
		return int64(val.Int16()), true
	case proto.TypeUint16:
		return int64(val.Uint16()), true
	case proto.TypeInt32:
		return int64(val.Int32()), true
	case proto.TypeUint32:
		return int64(val.Uint32()), true
	case proto.TypeInt64:
		return val.Int64(), true
	case proto.TypeUint64:
		return int64(val.Uint64()), true
	}
	return 0, false
}
