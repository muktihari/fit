// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"fmt"
	"unicode/utf8"

	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const (
	errNoFields                = errorString("no fields")
	errValueTypeMismatch       = errorString("value type mismatch")
	errInvalidUTF8String       = errorString("invalid UTF-8 string")
	errExceedMaxAllowed        = errorString("exceed max allowed")
	errMissingDeveloperDataId  = errorString("missing developer data id")
	errMissingFieldDescription = errorString("missing field description")
)

// MessageValidator is an interface for implementing message validation before encoding the message.
type MessageValidator interface {
	// Validate performs message validation before encoding to avoid resulting a corrupt FIT file.
	//
	// The validation process includes:
	//   - Removing fields created during component expansion.
	//   - Restoring float64-scaled field values to their binary forms (sint, uint, etc.).
	//   - Verifying whether the type and value are in alignment.
	Validate(mesg *proto.Message) error

	// Reset the message validator.
	Reset()
}

type messageValidator struct {
	developerDataIndexSeen [4]uint64 // 256-bits bitmap for checking if DeveloperDataIndex has been seen.
	fieldDescriptions      []*mesgdef.FieldDescription
}

func (v *messageValidator) Reset() {
	v.developerDataIndexSeen = [4]uint64{}
	clear(v.fieldDescriptions) // avoid memory leaks
	v.fieldDescriptions = v.fieldDescriptions[:0]
}

func (v *messageValidator) Validate(mesg *proto.Message) error {
	var valid int
	for i := range mesg.Fields {
		field := &mesg.Fields[i]

		if field.FieldBase == nil || field.IsExpandedField {
			continue
		}

		// Restore any scaled float64 value back into its corresponding integer representation.
		if field.Scale != 1 || field.Offset != 0 {
			field.Value = scaleoffset.DiscardValue(
				field.Value,
				field.BaseType,
				field.Scale,
				field.Offset,
			)
		}

		if err := valueIntegrity(field.Value, field.BaseType); err != nil {
			return fmt.Errorf("field index: %d, num: %d, name: %s: %w", i, field.Num, field.Name, err)
		}

		if i != valid {
			mesg.Fields[i], mesg.Fields[valid] = mesg.Fields[valid], mesg.Fields[i]
		}

		if valid == 255 {
			return fmt.Errorf("max n fields is 255: %w", errExceedMaxAllowed)
		}

		valid++
	}

	mesg.Fields = mesg.Fields[:valid]
	if len(mesg.Fields) == 0 && len(mesg.DeveloperFields) == 0 {
		return errNoFields
	}

	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		x := mesg.FieldValueByNum(fieldnum.DeveloperDataIdDeveloperDataIndex).Uint8()
		v.developerDataIndexSeen[x>>6] |= 1 << (x & 63)
	case mesgnum.FieldDescription:
		v.fieldDescriptions = append(v.fieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}

	if len(mesg.DeveloperFields) == 0 {
		return nil
	}

	for i := range mesg.DeveloperFields {
		developerField := &mesg.DeveloperFields[i]

		if x := developerField.DeveloperDataIndex; (v.developerDataIndexSeen[x>>6]>>(x&63))&1 == 0 {
			return fmt.Errorf("developer field index: %d, num: %d: %w",
				i, developerField.Num, errMissingDeveloperDataId)
		}

		var fieldDesc *mesgdef.FieldDescription
		for _, f := range v.fieldDescriptions {
			if f.DeveloperDataIndex == developerField.DeveloperDataIndex &&
				f.FieldDefinitionNumber == developerField.Num {
				fieldDesc = f
				break
			}
		}

		if fieldDesc == nil {
			return fmt.Errorf("developer field index: %d, num: %d: %w",
				i, developerField.Num, errMissingFieldDescription)
		}

		// Restore any scaled float64 value back into its corresponding integer representation.
		if fieldDesc.Scale != basetype.Uint8Invalid && fieldDesc.Offset != basetype.Sint8Invalid {
			developerField.Value = scaleoffset.DiscardValue(
				developerField.Value,
				fieldDesc.FitBaseTypeId,
				float64(fieldDesc.Scale),
				float64(fieldDesc.Offset),
			)
		}

		if err := valueIntegrity(developerField.Value, fieldDesc.FitBaseTypeId); err != nil {
			return fmt.Errorf("developer field index: %d, num: %d: %w",
				i, developerField.Num, err)
		}

		if i == 255 {
			return fmt.Errorf("max n developer fields is 255: %w", errExceedMaxAllowed)
		}
	}

	return nil
}

func valueIntegrity(value proto.Value, baseType basetype.BaseType) error {
	if !value.Align(baseType) {
		return fmt.Errorf("value %v with proto type '%v' is not align with the expected basetype '%s': %w",
			value, value.Type(), baseType, errValueTypeMismatch)
	}

	// UTF-8 String Validation
	switch value.Type() {
	case proto.TypeString:
		val := value.String()
		if !utf8.ValidString(val) {
			return fmt.Errorf("%q is not a valid utf-8 string: %w", val, errInvalidUTF8String)
		}
	case proto.TypeSliceString:
		val := value.SliceString()
		for i := range val {
			if !utf8.ValidString(val[i]) {
				return fmt.Errorf("[%d] %q is not a valid utf-8 string: %w", i, val[i], errInvalidUTF8String)
			}
		}
	}

	// Both proto.FieldDefinition's Size and proto.DeveloperFieldDefinition's Size is a type of byte.
	if size := value.Size(); size > 255 {
		return fmt.Errorf("max value size in bytes is 255, got: %d (value: %v): %w",
			size, value.Any(), errExceedMaxAllowed)
	}

	return nil
}
