// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"errors"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestMessageValidatorOption(t *testing.T) {
	tt := []struct {
		name    string
		opts    []ValidatorOption
		options *validatorOptions
	}{
		{
			name: "defaultValidatorOptions",
			options: &validatorOptions{
				omitInvalidValues: true,
			},
		},
		{
			name: "with options",
			opts: []ValidatorOption{
				ValidatorWithPreserveInvalidValues(),
			},
			options: &validatorOptions{
				omitInvalidValues: false,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mv := NewMessageValidator(tc.opts...).(*messageValidator)
			if diff := cmp.Diff(mv.options, tc.options, cmp.AllowUnexported(validatorOptions{})); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageValidatorValidate(t *testing.T) {
	tt := []struct {
		name          string
		mesgs         []proto.Message
		mesgValidator MessageValidator
		errs          []error
	}{
		{
			name: "valid message without developer fields happy flow",
			mesgs: []proto.Message{
				{
					Fields: []proto.Field{
						factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
					},
				},
			},
		},
		{
			name: "valid message with developer fields happy flow",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "Heart Rate",
					fieldnum.FieldDescriptionNativeMesgNum:         mesgnum.Record,
					fieldnum.FieldDescriptionNativeFieldNum:        fieldnum.RecordHeartRate,
					fieldnum.FieldDescriptionFitBaseTypeId:         basetype.Uint8,
				}),
				{
					Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
					},
					DeveloperFields: []proto.DeveloperField{
						{
							DeveloperDataIndex: 0,
							Num:                0,
							Size:               1,
							Name:               "Heart Rate",
							NativeMesgNum:      mesgnum.Record,
							NativeFieldNum:     fieldnum.RecordHeartRate,
							Type:               basetype.Uint8,
						},
					},
				},
			},
		},
		{
			name: "mesg contain expanded field",
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
					func() proto.Field {
						field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
						field.IsExpandedField = true
						field.Value = uint32(1000)
						return field
					}(),
				),
			},
		},
		{
			name: "mesg contain field with invalid size",
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.Record).WithFields(
					proto.Field{
						FieldBase: &proto.FieldBase{
							Size: 0,
						},
					},
				),
			},
			errs: []error{ErrSizeZero},
		},
		{
			name: "mesg contain field with scaled value",
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordAltitude).WithValue((float64(37304) / 5) - 500), // 6960.8m
				),
			},
		},
		{
			name: "mesg contain field value type not align",
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint32(1000)), // should be uint16
				),
			},
			errs: []error{ErrValueTypeMismatch},
		},
		{
			name: "valid message with developer data index not found in previous message sequence",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
				}),
				{
					Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
					},
					DeveloperFields: []proto.DeveloperField{
						{DeveloperDataIndex: 0, Num: 0},
					},
				},
			},
			errs: []error{nil, ErrMissingDeveloperDataId},
		},
		{
			name: "valid message with field description not found in previous message sequence",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				{
					Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
					},
					DeveloperFields: []proto.DeveloperField{
						{DeveloperDataIndex: 0, Num: 0},
					},
				},
			},
			errs: []error{nil, ErrMissingFieldDescription},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mesgValidator := tc.mesgValidator
			if mesgValidator == nil {
				mesgValidator = NewMessageValidator()
			}

			if tc.errs == nil {
				tc.errs = make([]error, len(tc.mesgs))
			}

			for i, mesg := range tc.mesgs {
				err := mesgValidator.Validate(&mesg)
				if !errors.Is(err, tc.errs[i]) {
					t.Fatalf("expected err: %v, got: %v", tc.errs[i], err)
				}
			}
		})
	}
}

func TestIsValueTypeAligned(t *testing.T) {
	type DefinedByte byte
	var i8 int8 = 10

	tt := []struct {
		value    any
		baseType basetype.BaseType
		expected bool
	}{
		{value: nil, baseType: basetype.Sint8, expected: false},
		{value: true, baseType: basetype.Enum, expected: true},
		{value: []bool{true, false}, baseType: basetype.Enum, expected: true},
		{value: &i8, baseType: basetype.Sint8, expected: true},
		{value: int8(1), baseType: basetype.Sint8, expected: true},
		{value: uint8(1), baseType: basetype.Uint8, expected: true},
		{value: uint8(1), baseType: basetype.Uint8z, expected: true},
		{value: int16(1), baseType: basetype.Sint16, expected: true},
		{value: uint16(1), baseType: basetype.Uint16, expected: true},
		{value: uint16(1), baseType: basetype.Uint16z, expected: true},
		{value: int32(1), baseType: basetype.Sint32, expected: true},
		{value: uint32(1), baseType: basetype.Uint32, expected: true},
		{value: uint32(1), baseType: basetype.Uint32z, expected: true},
		{value: float32(1.0), baseType: basetype.Float32, expected: true},
		{value: float64(1.0), baseType: basetype.Float64, expected: true},
		{value: int64(1.0), baseType: basetype.Sint64, expected: true},
		{value: uint64(1), baseType: basetype.Uint64, expected: true},
		{value: uint64(1), baseType: basetype.Uint64z, expected: true},
		{value: string("Fit SDK"), baseType: basetype.String, expected: true},
		{value: []byte("Fit SDK"), baseType: basetype.Byte, expected: true},
		{value: []any{byte(1), byte(2)}, baseType: basetype.Byte, expected: true},
		{value: []any{byte(1), DefinedByte(2)}, baseType: basetype.Byte, expected: true},
		{value: []any{DefinedByte(1), DefinedByte(2)}, baseType: basetype.Byte, expected: true},
		{value: []any{DefinedByte(1), int8(2)}, baseType: basetype.Byte, expected: false},
		{value: []*int8{&i8}, baseType: basetype.Sint8, expected: false},
		{value: []int8{1, 2, 3}, baseType: basetype.Sint8, expected: true},
		{value: DefinedByte(1), baseType: basetype.Byte, expected: true},
		{value: DefinedByte(1), baseType: basetype.Sint16, expected: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v (%T): %s", tc.value, tc.value, tc.baseType), func(t *testing.T) {
			if isAlign := isValueTypeAligned(tc.value, tc.baseType); isAlign != tc.expected {
				t.Fatalf("expected: %t, got %t", tc.expected, isAlign)
			}
		})
	}
}

func TestHasValidValue(t *testing.T) {
	type Int8 int8
	type Uint8 uint8
	type Int16 int16
	type Uint16 uint16
	type Int32 int32
	type Uint32 uint32
	type String string
	type Float32 float32
	type Float64 float64
	type Int64 int64
	type Uint64 uint64

	tt := []struct {
		value    any
		expected bool
	}{
		{value: nil, expected: false},
		{value: int8(0), expected: true},
		{value: uint8(0), expected: true},
		{value: int16(0), expected: true},
		{value: uint16(0), expected: true},
		{value: int32(0), expected: true},
		{value: uint32(0), expected: true},
		{value: string("Fit SDK Go"), expected: true},
		{value: string(""), expected: false},
		{value: string("\x00"), expected: false},
		{value: float32(0.2), expected: true},
		{value: float32(math.Float32frombits(basetype.Float32Invalid)), expected: false},
		{value: float32(math.Float32frombits(basetype.Float32Invalid - 1)), expected: true},
		{value: float64(0.5), expected: true},
		{value: float64(math.Float64frombits(basetype.Float64Invalid)), expected: false},
		{value: float64(math.Float64frombits(basetype.Float64Invalid - 1)), expected: true},
		{value: int64(0), expected: true},
		{value: uint64(0), expected: true},
		{value: struct{}{}, expected: false},
		{value: []int8{0, basetype.Sint8Invalid}, expected: true},
		{value: []uint8{0, basetype.Uint8Invalid}, expected: true},
		{value: []int16{0, basetype.Sint16Invalid}, expected: true},
		{value: []uint16{0, basetype.Uint16Invalid}, expected: true},
		{value: []int32{0, basetype.Sint32Invalid}, expected: true},
		{value: []string{"Fit SDK Go"}, expected: true},
		{value: []string{""}, expected: false},
		{value: []string{"\x00"}, expected: false},
		{value: []uint32{0, basetype.Uint32Invalid}, expected: true},
		{value: []float32{0.2, math.Float32frombits(basetype.Float32Invalid)}, expected: true},
		{value: []float64{0.5, math.Float64frombits(basetype.Float64Invalid)}, expected: true},
		{value: []int64{0, basetype.Sint64Invalid}, expected: true},
		{value: []uint64{0, basetype.Uint64Invalid}, expected: true},
		{value: Int8(0), expected: true},
		{value: Uint8(0), expected: true},
		{value: Int16(0), expected: true},
		{value: Uint16(0), expected: true},
		{value: Int32(0), expected: true},
		{value: Uint32(0), expected: true},
		{value: String("Fit SDK Go"), expected: true},
		{value: Float32(0.2), expected: true},
		{value: Float64(0.5), expected: true},
		{value: Int64(0), expected: true},
		{value: Uint64(0), expected: true},
		{value: []Int8{0, Int8(basetype.Sint8Invalid)}, expected: true},
		{value: []Uint8{0, Uint8(basetype.Uint8Invalid)}, expected: true},
		{value: []Int16{0, Int16(basetype.Sint16Invalid)}, expected: true},
		{value: []Uint16{0, Uint16(basetype.Uint16Invalid)}, expected: true},
		{value: []Int32{0, Int32(basetype.Sint32Invalid)}, expected: true},
		{value: []Uint32{0, Uint32(basetype.Uint32Invalid)}, expected: true},
		{value: []String{"Fit SDK Go"}, expected: true},
		{value: []String{""}, expected: false},
		{value: []String{"\x00"}, expected: false},
		{value: []Float32{0.2, Float32(math.Float32frombits(basetype.Float32Invalid))}, expected: true},
		{value: []Float64{0.5, Float64(math.Float64frombits(basetype.Float64Invalid))}, expected: true},
		{value: []Int64{0, Int64(basetype.Sint64Invalid)}, expected: true},
		{value: []Uint64{0, Uint64(basetype.Uint64Invalid)}, expected: true},
		{value: []struct{}{{}}, expected: false},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v (%T)", tc.value, tc.value), func(t *testing.T) {
			result := hasValidValue(tc.value)
			if result != tc.expected {
				t.Fatalf("expected: %t, got: %t", tc.expected, result)
			}
		})
	}
}

func BenchmarkIsValueTypeAligned(b *testing.B) {
	b.Run("benchmark primitive-values byte", func(b *testing.B) {
		var value byte = 10
		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Byte)
			if !ok {
				b.Fail()
			}
		}
	})

	b.Run("benchmark primitive-values float64", func(b *testing.B) {
		var value float64 = 10.5
		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Float64)
			if !ok {
				b.Fail()
			}
		}
	})

	b.Run("benchmark types", func(b *testing.B) {
		var value typedef.File = typedef.FileActivity
		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Byte)
			if !ok {
				b.Fail()
			}
		}
	})

	b.Run("benchmark []byte", func(b *testing.B) {
		var value = []byte{1, 2, 3, 4, 5}

		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Byte)
			if !ok {
				b.Fail()
			}
		}
	})

	b.Run("benchmark []float64", func(b *testing.B) {
		var value = []float64{
			1.9123455,
			2.9123455,
			3.9123455,
			4.9123455,
			5.9123455,
		}

		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Float64)
			if !ok {
				b.Fail()
			}
		}
	})

	b.Run("benchmark []typedef.File", func(b *testing.B) {
		var value = []typedef.File{
			typedef.FileActivity,
			typedef.FileActivitySummary,
			typedef.FileCourse,
			typedef.FileBloodPressure,
			typedef.FileDevice,
		}

		for i := 0; i < b.N; i++ {
			ok := isValueTypeAligned(value, basetype.Byte)
			if !ok {
				b.Fail()
			}
		}
	})
}

func BenchmarkValidate(b *testing.B) {
	b.StopTimer()
	mesgValidator := NewMessageValidator()
	mesg := factory.CreateMesgOnly(mesgnum.Record).WithFields(
		factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
		factory.CreateField(mesgnum.Record, fieldnum.RecordAltitude).WithValue(uint16(10000)),
		func() proto.Field {
			field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
			field.IsExpandedField = true
			field.Value = uint32(1000)
			return field
		}(),
		func() proto.Field {
			field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
			field.IsExpandedField = true
			field.Value = uint32(1000)
			return field
		}(),
		func() proto.Field {
			field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
			field.IsExpandedField = true
			field.Value = uint32(1000)
			return field
		}(),
	)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = mesgValidator.Validate(&mesg)
	}
}
