// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"fmt"
	"unicode/utf8"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const (
	ErrInvalidUTF8String       = errorString("invalid UTF-8 string")
	ErrValueTypeMismatch       = errorString("value type mismatch")
	ErrNoFields                = errorString("no fields")
	ErrMissingDeveloperDataId  = errorString("missing developer data id")
	ErrMissingFieldDescription = errorString("missing field description")
	ErrExceedMaxAllowed        = errorString("exceed max allowed")
)

// MessageValidator is an interface for implementing message validation before encoding the message.
type MessageValidator interface {
	// Validate performs message validation before encoding to avoid resulting a corrupt FIT file.
	//
	// The validation process includes:
	//   1. Removing fields created during component expansion.
	//   2. Removing fields with invalid values.
	//   3. Restoring float64-scaled field values to their binary forms (sint, uint, etc.).
	//   4. Verifying whether the type and value are in alignment.
	Validate(mesg *proto.Message) error

	// Reset the message validator.
	Reset()
}

type validatorOptions struct {
	omitInvalidValues bool // default: true -> field containing invalid values will be omitted.
	factory           Factory
}

// ValidatorOption is message validator's option.
type ValidatorOption func(o *validatorOptions)

func defaultValidatorOptions() validatorOptions {
	return validatorOptions{
		omitInvalidValues: true,
		factory:           factory.StandardFactory(),
	}
}

// Factory defines a contract that any Factory containing these method can be used by the Encoder's Validator.
type Factory interface {
	// CreateField creates new field based on defined messages in the factory.
	// If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

// ValidatorWithPreserveInvalidValues directs the message validator to preserve invalid value instead of omit it.
func ValidatorWithPreserveInvalidValues() ValidatorOption {
	return func(o *validatorOptions) {
		o.omitInvalidValues = false
	}
}

// ValidatorWithFactory directs the message validator to use this factory instead of standard factory.
// The factory is only used for validating developer fields that have valid native data.
func ValidatorWithFactory(factory Factory) ValidatorOption {
	return func(o *validatorOptions) {
		if o.factory != nil {
			o.factory = factory
		}
	}
}

type messageValidator struct {
	options              validatorOptions
	developerDataIndexes []uint8
	fieldDescriptions    []*mesgdef.FieldDescription
}

// NewMessageValidator creates new message validator. The validator is mainly used to validate message before encoding.
// This receives options that direct the message validator how it should behave in certain way.
func NewMessageValidator(opts ...ValidatorOption) MessageValidator {
	mv := &messageValidator{}
	mv.options = defaultValidatorOptions()
	for i := range opts {
		opts[i](&mv.options)
	}
	return mv
}

func (v *messageValidator) Reset() {
	v.developerDataIndexes = v.developerDataIndexes[:0]

	for i := range v.fieldDescriptions {
		v.fieldDescriptions[i] = nil // avoid memory leaks
	}
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

		if v.options.omitInvalidValues && !field.Value.Valid(field.BaseType) {
			continue
		}

		if err := valueIntegrity(field.Value, field.BaseType); err != nil {
			return fmt.Errorf("field index: %d, num: %d, name: %s: %w", i, field.Num, field.Name, err)
		}

		if i != valid {
			mesg.Fields[i], mesg.Fields[valid] = mesg.Fields[valid], mesg.Fields[i]
		}
		if valid == 255 {
			return fmt.Errorf("max n fields is 255: %w", ErrExceedMaxAllowed)
		}
		valid++
	}

	mesg.Fields = mesg.Fields[:valid]
	if len(mesg.Fields) == 0 && len(mesg.DeveloperFields) == 0 {
		return ErrNoFields
	}

	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		v.developerDataIndexes = append(v.developerDataIndexes,
			mesg.FieldValueByNum(fieldnum.DeveloperDataIdDeveloperDataIndex).Uint8())
	case mesgnum.FieldDescription:
		v.fieldDescriptions = append(v.fieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}

	if len(mesg.DeveloperFields) == 0 {
		return nil
	}

	valid = 0
	for i := range mesg.DeveloperFields {
		developerField := &mesg.DeveloperFields[i]

		var ok bool
		for _, developerDataIndex := range v.developerDataIndexes {
			if developerDataIndex == developerField.DeveloperDataIndex {
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("developer field index: %d, num: %d: %w",
				i, developerField.Num, ErrMissingDeveloperDataId)
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
				i, developerField.Num, ErrMissingFieldDescription)
		}

		// Restore any scaled float64 value back into its corresponding integer representation.
		if fieldDesc.NativeMesgNum != typedef.MesgNumInvalid && fieldDesc.NativeFieldNum != basetype.Uint8Invalid {
			// If Developer Field contains a valid NativeMesgNum and NativeFieldNum,
			// the value should be treated as native value (scale, offset, etc shall apply).
			field := v.options.factory.CreateField(fieldDesc.NativeMesgNum, fieldDesc.NativeFieldNum)
			if field.Name != factory.NameUnknown && (field.Scale != 1 || field.Offset != 0) {
				developerField.Value = scaleoffset.DiscardValue(
					developerField.Value,
					field.BaseType,
					field.Scale,
					field.Offset,
				)
			}
		} else if fieldDesc.Scale != basetype.Uint8Invalid && fieldDesc.Offset != basetype.Sint8Invalid {
			developerField.Value = scaleoffset.DiscardValue(
				developerField.Value,
				fieldDesc.FitBaseTypeId,
				float64(fieldDesc.Scale),
				float64(fieldDesc.Offset),
			)
		}

		if v.options.omitInvalidValues && !developerField.Value.Valid(fieldDesc.FitBaseTypeId) {
			continue
		}

		if err := valueIntegrity(developerField.Value, fieldDesc.FitBaseTypeId); err != nil {
			return fmt.Errorf("developer field index: %d, num: %d: %w",
				i, developerField.Num, err)
		}

		if i != valid {
			mesg.DeveloperFields[i], mesg.DeveloperFields[valid] = mesg.DeveloperFields[valid], mesg.DeveloperFields[i]
		}
		if valid == 255 {
			return fmt.Errorf("max n developer fields is 255: %w", ErrExceedMaxAllowed)
		}
		valid++
	}

	mesg.DeveloperFields = mesg.DeveloperFields[:valid]

	return nil
}

func valueIntegrity(value proto.Value, baseType basetype.BaseType) error {
	if !value.Align(baseType) {
		val := value.Any()
		return fmt.Errorf("value %v with type '%T' is not align with the expected type '%s': %w",
			val, val, baseType, ErrValueTypeMismatch)
	}

	// UTF-8 String Validation
	switch value.Type() {
	case proto.TypeString:
		val := value.String()
		if !utf8.ValidString(val) {
			return fmt.Errorf("%q is not a valid utf-8 string: %w", val, ErrInvalidUTF8String)
		}
	case proto.TypeSliceString:
		val := value.SliceString()
		for i := range val {
			if !utf8.ValidString(val[i]) {
				return fmt.Errorf("[%d] %q is not a valid utf-8 string: %w", i, val[i], ErrInvalidUTF8String)
			}
		}
	}

	// Both proto.FieldDefinition's Size and proto.DeveloperFieldDefinition's Size is a type of byte.
	size := proto.Sizeof(value)
	if size > 255 {
		return fmt.Errorf("max value size in bytes is 255, got: %d: %w", size, ErrExceedMaxAllowed)
	}

	return nil
}
