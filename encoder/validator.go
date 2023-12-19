// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	ErrSizeZero                = errors.New("size is zero")
	ErrValueTypeMismatch       = errors.New("value type mismatch")
	ErrNoFields                = errors.New("no fields")
	ErrMissingDeveloperDataId  = errors.New("missing developer data id")
	ErrMissingFieldDescription = errors.New("missing field description")
)

// MessageValidator is an interface for implementing message validation before encoding the message.
type MessageValidator interface {
	// Validate performs message validation before encoding to avoid resulting a corrupt Fit file.
	//
	// The validation process includes:
	//   1. Removing fields created during component expansion.
	//   2. Removing fields with invalid values.
	//   3. Restoring float64-scaled field values to their binary forms (sint, uint, etc.).
	//   4. Verifying whether the type and value are in alignment.
	Validate(mesg *proto.Message) error
}

type validatorOptions struct {
	omitInvalidValues bool // default: true -> field containing invalid values will be omitted.
}

type ValidatorOption interface{ apply(o *validatorOptions) }

type fnApplyValidatorOption func(o *validatorOptions)

func (f fnApplyValidatorOption) apply(o *validatorOptions) { f(o) }

func defaultValidatorOptions() *validatorOptions {
	return &validatorOptions{
		omitInvalidValues: true,
	}
}

func ValidatorWithPreserveInvalidValues() ValidatorOption {
	return fnApplyValidatorOption(func(o *validatorOptions) {
		o.omitInvalidValues = false
	})
}

func NewMessageValidator(opts ...ValidatorOption) MessageValidator {
	options := defaultValidatorOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	return &messageValidator{
		options: options,
	}
}

type messageValidator struct {
	options           *validatorOptions
	developerDataIds  []*mesgdef.DeveloperDataId
	fieldDescriptions []*mesgdef.FieldDescription
}

func (v *messageValidator) Validate(mesg *proto.Message) error {
	mesg.Header = proto.MesgNormalHeaderMask // reset default

	fields := make([]proto.Field, 0, v.size(mesg.Fields))
	for i := range mesg.Fields {
		field := &mesg.Fields[i]

		if field.FieldBase == nil || field.IsExpandedField {
			continue
		}

		if field.Size == 0 {
			return fmt.Errorf("size could not be zero for fieldIndex: %d fieldNum: %d, fieldName: %s: %w",
				i, field.Num, field.Name, ErrSizeZero)
		}

		if v.options.omitInvalidValues && !hasValidValue(field.Value) {
			continue
		}

		// Restore any scaled float64 value back into its corresponding integer representation.
		if field.Scale != 1 && field.Offset != 0 {
			field.Value = scaleoffset.DiscardAny(
				field.Value,
				field.Type.BaseType(),
				field.Scale,
				field.Offset,
			)
		}

		// Now that the value is sanitized, we can check whether the type and value are aligned.
		if !isValueTypeAligned(field.Value, field.Type.BaseType()) {
			return fmt.Errorf(
				"type '%T' is not align with the expected type '%s' for fieldIndex: %d, fieldNum: %d, fieldName: %q, fieldValue: '%v': %w",
				field.Value, field.Type, i, field.Num, field.Name, field.Value, ErrValueTypeMismatch)
		}

		fields = append(fields, *field)
	}

	if len(fields) == 0 && len(mesg.DeveloperFields) == 0 {
		return ErrNoFields
	}

	mesg.Fields = fields

	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		v.developerDataIds = append(v.developerDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		v.fieldDescriptions = append(v.fieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}

	if len(mesg.DeveloperFields) == 0 {
		return nil
	}

	for i := range mesg.DeveloperFields {
		devField := &mesg.DeveloperFields[i]

		var isDeveloperDataIndexPresent bool
		for _, developerDataId := range v.developerDataIds {
			if devField.DeveloperDataIndex == developerDataId.DeveloperDataIndex {
				isDeveloperDataIndexPresent = true
				break
			}
		}

		if !isDeveloperDataIndexPresent {
			return fmt.Errorf("developer field index: %d, num: %d: %w", i, devField.Num, ErrMissingDeveloperDataId)
		}

		var isFieldDescriptionPresent bool
		for _, fieldDescription := range v.fieldDescriptions {
			if devField.DeveloperDataIndex == fieldDescription.DeveloperDataIndex &&
				devField.Num == fieldDescription.FieldDefinitionNumber {
				isFieldDescriptionPresent = true
				break
			}
		}

		if !isFieldDescriptionPresent {
			return fmt.Errorf("developer field index: %d, num: %d: %w", i, devField.Num, ErrMissingFieldDescription)
		}
	}

	return nil
}

// size return the size of valid fields.
// This is method ensure we just do 1 alloc on creating new slice of fields.
func (v *messageValidator) size(fields []proto.Field) byte {
	var size byte
	for i := range fields {
		field := &fields[i]
		if field.FieldBase == nil || field.IsExpandedField {
			continue
		}
		if v.options.omitInvalidValues && !hasValidValue(field.Value) {
			continue
		}
		size++
	}
	return size
}

// isValueTypeAligned checks whether the value is aligned with type. The value should be a concrete type not pointer to a value.
func isValueTypeAligned(value any, baseType basetype.BaseType) bool {
	if value == nil {
		return false
	}

	switch val := value.(type) { // Fast path
	case bool, []bool:
		return baseType == basetype.Enum
	case int8, []int8:
		return baseType == basetype.Sint8
	case uint8, []uint8: // == byte
		return baseType == basetype.Enum ||
			baseType == basetype.Byte ||
			baseType == basetype.Uint8 ||
			baseType == basetype.Uint8z
	case int16, []int16:
		return baseType == basetype.Sint16
	case uint16, []uint16:
		return baseType == basetype.Uint16 || baseType == basetype.Uint16z
	case int32, []int32:
		return baseType == basetype.Sint32
	case uint32, []uint32:
		return baseType == basetype.Uint32 || baseType == basetype.Uint32z
	case float32, []float32:
		return baseType == basetype.Float32
	case float64, []float64:
		return baseType == basetype.Float64
	case int64, []int64:
		return baseType == basetype.Sint64
	case uint64, []uint64:
		return baseType == basetype.Uint64 || baseType == basetype.Uint64z
	case string: // we have no []string values
		return baseType == basetype.String
	case []any:
		for i := range val {
			if !isValueTypeAligned(val[i], baseType) {
				return false
			}
		}
		return true
	}

	// Fallback to reflection, reflect.TypeOf is fast enough and manageable. In this case, it's preferrable than asserting every possible types.
	rt := reflect.TypeOf(value)
	switch rt.Kind() {
	case reflect.Pointer, reflect.Slice:
		return rt.Elem().Kind() == baseType.Kind()
	}

	return rt.Kind() == baseType.Kind()
}

// hasValidValue checks whether given val has any valid value.
// If val is a slice, even though only one value is valid, it will be counted a valid value.
func hasValidValue(val any) bool {
	if val == nil {
		return false
	}

	switch v := val.(type) { // Fast Path
	case int8:
		return v != basetype.Sint8Invalid
	case uint8:
		return v != basetype.Uint8Invalid
	case int16:
		return v != basetype.Sint16Invalid
	case uint16:
		return v != basetype.Uint16Invalid
	case int32:
		return v != basetype.Sint32Invalid
	case uint32:
		return v != basetype.Uint32Invalid
	case string:
		return v != basetype.StringInvalid && v != ""
	case float32:
		return math.Float32bits(v) != basetype.Float32Invalid
	case float64:
		return math.Float64bits(v) != basetype.Float64Invalid
	case int64:
		return v != basetype.Sint64Invalid
	case uint64:
		return v != basetype.Uint64Invalid
	case []int8:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Sint8Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []uint8:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Uint8Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []int16:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Sint16Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []uint16:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Uint16Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []int32:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Sint32Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []uint32:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Uint32Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []string:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.StringInvalid || v[i] == "" {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []float32:
		invalidcounter := 0
		for i := range v {
			if math.Float32bits(v[i]) == basetype.Float32Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []float64:
		invalidcounter := 0
		for i := range v {
			if math.Float64bits(v[i]) == basetype.Float64Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []int64:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Sint64Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	case []uint64:
		invalidcounter := 0
		for i := range v {
			if v[i] == basetype.Uint64Invalid {
				invalidcounter++
			}
		}
		return invalidcounter != len(v)
	}

	// Fallback to reflection
	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Int8:
		return int8(rv.Int()) != basetype.Sint8Invalid
	case reflect.Uint8:
		return uint8(rv.Uint()) != basetype.Uint8Invalid
	case reflect.Int16:
		return int16(rv.Int()) != basetype.Sint16Invalid
	case reflect.Uint16:
		return uint16(rv.Uint()) != basetype.Uint16Invalid
	case reflect.Int32:
		return int32(rv.Int()) != basetype.Sint32Invalid
	case reflect.Uint32:
		return uint32(rv.Uint()) != basetype.Uint32Invalid
	case reflect.String:
		return rv.String() != basetype.StringInvalid
	case reflect.Float32:
		return math.Float32bits(float32(rv.Float())) != basetype.Float32Invalid
	case reflect.Float64:
		return math.Float64bits(rv.Float()) != basetype.Float64Invalid
	case reflect.Int64:
		return int64(rv.Int()) != basetype.Sint64Invalid
	case reflect.Uint64:
		return uint64(rv.Uint()) != basetype.Uint64Invalid
	case reflect.Slice:
		invalidcounter := 0
		for i := 0; i < rv.Len(); i++ {
			rve := rv.Index(i)

			switch rve.Kind() {
			case reflect.Int8:
				if int8(rve.Int()) == basetype.Sint8Invalid {
					invalidcounter++
				}
			case reflect.Uint8:
				if uint8(rve.Uint()) == basetype.Uint8Invalid {
					invalidcounter++
				}
			case reflect.Int16:
				if int16(rve.Int()) == basetype.Sint16Invalid {
					invalidcounter++
				}
			case reflect.Uint16:
				if uint16(rve.Uint()) == basetype.Uint16Invalid {
					invalidcounter++
				}
			case reflect.Int32:
				if int32(rve.Int()) == basetype.Sint32Invalid {
					invalidcounter++
				}
			case reflect.Uint32:
				if uint32(rve.Uint()) == basetype.Uint32Invalid {
					invalidcounter++
				}
			case reflect.String:
				if str := rve.String(); str == basetype.StringInvalid || str == "" {
					invalidcounter++
				}
			case reflect.Float32:
				if math.Float32bits(float32(rve.Float())) == basetype.Float32Invalid {
					invalidcounter++
				}
			case reflect.Float64:
				if math.Float64bits(rve.Float()) == basetype.Float64Invalid {
					invalidcounter++
				}
			case reflect.Int64:
				if int64(rve.Int()) == basetype.Sint64Invalid {
					invalidcounter++
				}
			case reflect.Uint64:
				if uint64(rve.Uint()) == basetype.Uint64Invalid {
					invalidcounter++
				}
			default: // not supported
				return false
			}
		}
		return invalidcounter != rv.Len()
	default: // not supported
		return false
	}
}
