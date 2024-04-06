// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	ErrInvalidUTF8String       = errors.New("invalid UTF-8 string")
	ErrValueTypeMismatch       = errors.New("value type mismatch")
	ErrNoFields                = errors.New("no fields")
	ErrMissingDeveloperDataId  = errors.New("missing developer data id")
	ErrMissingFieldDescription = errors.New("missing field description")
	ErrExceedMaxAllowed        = errors.New("exceed max allowed")
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

type ValidatorOption interface{ apply(o *validatorOptions) }

type fnApplyValidatorOption func(o *validatorOptions)

func (f fnApplyValidatorOption) apply(o *validatorOptions) { f(o) }

func defaultValidatorOptions() *validatorOptions {
	return &validatorOptions{
		omitInvalidValues: true,
		factory:           factory.StandardFactory(),
	}
}

// Factory defines a contract that any Factory containing these method can be used by the Encoder's Validator.
type Factory interface {
	// CreateField create new field based on defined messages in the factory. If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

func ValidatorWithPreserveInvalidValues() ValidatorOption {
	return fnApplyValidatorOption(func(o *validatorOptions) {
		o.omitInvalidValues = false
	})
}

func ValidatorWithFactory(factory Factory) ValidatorOption {
	return fnApplyValidatorOption(func(o *validatorOptions) {
		if o.factory != nil {
			o.factory = factory
		}
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

	var valid byte
	for i := 0; i < len(mesg.Fields); i++ {
		field := &mesg.Fields[i]

		if field.FieldBase == nil || field.IsExpandedField {
			continue
		}

		if v.options.omitInvalidValues && !hasValidValue(field.Value) {
			continue
		}

		// Restore any scaled float64 value back into its corresponding integer representation.
		if field.Scale != 1 && field.Offset != 0 {
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

		mesg.Fields[i], mesg.Fields[valid] = mesg.Fields[valid], mesg.Fields[i]
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
		v.developerDataIds = append(v.developerDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		v.fieldDescriptions = append(v.fieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}

	if len(mesg.DeveloperFields) == 0 {
		return nil
	}

	valid = 0
	for i := range mesg.DeveloperFields {
		developerField := &mesg.DeveloperFields[i]

		if err := v.validateReference(developerField); err != nil {
			return fmt.Errorf("developer field index: %d, num: %d, name: %s: %w",
				i, developerField.Num, developerField.Name, err)
		}

		if v.options.omitInvalidValues && !hasValidValue(developerField.Value) {
			continue
		}

		v.handleNativeValue(developerField)

		if err := valueIntegrity(developerField.Value, developerField.BaseType); err != nil {
			return fmt.Errorf("developer field index: %d, num: %d, name: %s: %w",
				i, developerField.Num, developerField.Name, err)
		}

		mesg.DeveloperFields[i], mesg.DeveloperFields[valid] = mesg.DeveloperFields[valid], mesg.DeveloperFields[i]
		if valid == 255 {
			return fmt.Errorf("max n developer fields is 255: %w", ErrExceedMaxAllowed)
		}
		valid++
	}

	mesg.DeveloperFields = mesg.DeveloperFields[:valid]

	return nil
}

func (v *messageValidator) validateReference(developerField *proto.DeveloperField) error {
	var ok bool
	for _, d := range v.developerDataIds {
		if developerField.DeveloperDataIndex == d.DeveloperDataIndex {
			ok = true
			break
		}
	}
	if !ok {
		return ErrMissingDeveloperDataId
	}

	for _, f := range v.fieldDescriptions {
		if developerField.DeveloperDataIndex == f.DeveloperDataIndex &&
			developerField.Num == f.FieldDefinitionNumber {
			return nil
		}
	}
	return ErrMissingFieldDescription
}

// If Developer Field contains a valid NativeMesgNum and NativeFieldNum,
// the value should be treated as native value (scale, offset, etc shall apply).
func (v *messageValidator) handleNativeValue(developerField *proto.DeveloperField) {
	if developerField.NativeMesgNum == 0 && developerField.NativeFieldNum == 0 {
		return
	}

	field := v.options.factory.CreateField(developerField.NativeMesgNum, developerField.NativeFieldNum)
	if field.Name == factory.NameUnknown { // Unknown Field will always have Scale: 1 and Offset: 0.
		return
	}

	// Restore any scaled float64 value back into its corresponding integer representation.
	if field.Scale != 1 && field.Offset != 0 {
		developerField.Value = scaleoffset.DiscardValue(
			developerField.Value,
			field.BaseType,
			field.Scale,
			field.Offset,
		)
	}
}

func (v *messageValidator) Reset() {
	v.developerDataIds = v.developerDataIds[:0]
	v.fieldDescriptions = v.fieldDescriptions[:0]
}

func valueIntegrity(value proto.Value, baseType basetype.BaseType) error {
	if !isValueTypeAligned(value, baseType) {
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
	size := proto.Sizeof(value, baseType)
	if size > 255 {
		return fmt.Errorf("max value size in bytes is 255, got: %d: %w", size, ErrExceedMaxAllowed)
	}

	return nil
}

// isValueTypeAligned checks whether the value is aligned with type.
func isValueTypeAligned(value proto.Value, baseType basetype.BaseType) bool {
	switch value.Type() {
	case proto.TypeBool, proto.TypeSliceBool:
		return baseType == basetype.Enum
	case proto.TypeInt8, proto.TypeSliceInt8:
		return baseType == basetype.Sint8
	case proto.TypeUint8, proto.TypeSliceUint8: // == byte
		return baseType == basetype.Enum ||
			baseType == basetype.Byte ||
			baseType == basetype.Uint8 ||
			baseType == basetype.Uint8z
	case proto.TypeInt16, proto.TypeSliceInt16:
		return baseType == basetype.Sint16
	case proto.TypeUint16, proto.TypeSliceUint16:
		return baseType == basetype.Uint16 || baseType == basetype.Uint16z
	case proto.TypeInt32, proto.TypeSliceInt32:
		return baseType == basetype.Sint32
	case proto.TypeUint32, proto.TypeSliceUint32:
		return baseType == basetype.Uint32 || baseType == basetype.Uint32z
	case proto.TypeFloat32, proto.TypeSliceFloat32:
		return baseType == basetype.Float32
	case proto.TypeFloat64, proto.TypeSliceFloat64:
		return baseType == basetype.Float64
	case proto.TypeInt64, proto.TypeSliceInt64:
		return baseType == basetype.Sint64
	case proto.TypeUint64, proto.TypeSliceUint64:
		return baseType == basetype.Uint64 || baseType == basetype.Uint64z
	case proto.TypeString, proto.TypeSliceString:
		return baseType == basetype.String
	}
	return false
}

// hasValidValue checks whether given val has any valid value.
// If val is a slice, even though only one value is valid, it will be counted a valid value.
//
// Special case: bool or slice of bool will always be valid since bool type is often used as a flag and
// there are only two possibility (true/false).
func hasValidValue(val proto.Value) bool {
	var invalidCount int

	switch val.Type() {
	case proto.TypeBool, proto.TypeSliceBool:
		return true // Mark as valid
	case proto.TypeInt8:
		return val.Int8() != basetype.Sint8Invalid
	case proto.TypeUint8:
		return val.Uint8() != basetype.Uint8Invalid
	case proto.TypeInt16:
		return val.Int16() != basetype.Sint16Invalid
	case proto.TypeUint16:
		return val.Uint16() != basetype.Uint16Invalid
	case proto.TypeInt32:
		return val.Int32() != basetype.Sint32Invalid
	case proto.TypeUint32:
		return val.Uint32() != basetype.Uint32Invalid
	case proto.TypeString:
		s := val.String()
		return s != basetype.StringInvalid && s != ""
	case proto.TypeFloat32:
		return math.Float32bits(val.Float32()) != basetype.Float32Invalid
	case proto.TypeFloat64:
		return math.Float64bits(val.Float64()) != basetype.Float64Invalid
	case proto.TypeInt64:
		return val.Int64() != basetype.Sint64Invalid
	case proto.TypeUint64:
		return val.Uint64() != basetype.Uint64Invalid
	case proto.TypeSliceInt8:
		vals := val.SliceInt8()
		for i := range vals {
			if vals[i] == basetype.Sint8Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceUint8:
		vals := val.SliceUint8()
		for i := range vals {
			if vals[i] == basetype.Uint8Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceInt16:
		vals := val.SliceInt16()
		for i := range vals {
			if vals[i] == basetype.Sint16Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceUint16:
		vals := val.SliceUint16()
		for i := range vals {
			if vals[i] == basetype.Uint16Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceInt32:
		vals := val.SliceInt32()
		for i := range vals {
			if vals[i] == basetype.Sint32Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceUint32:
		vals := val.SliceUint32()
		for i := range vals {
			if vals[i] == basetype.Uint32Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceString:
		vals := val.SliceString()
		for i := range vals {
			if vals[i] == basetype.StringInvalid || vals[i] == "" {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceFloat32:
		vals := val.SliceFloat32()
		for i := range vals {
			if math.Float32bits(vals[i]) == basetype.Float32Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceFloat64:
		vals := val.SliceFloat64()
		for i := range vals {
			if math.Float64bits(vals[i]) == basetype.Float64Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceInt64:
		vals := val.SliceInt64()
		for i := range vals {
			if vals[i] == basetype.Sint64Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case proto.TypeSliceUint64:
		vals := val.SliceUint64()
		for i := range vals {
			if vals[i] == basetype.Uint64Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	}

	return false
}
