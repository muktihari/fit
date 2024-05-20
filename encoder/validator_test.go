// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestMessageValidatorOption(t *testing.T) {
	fac := factory.New()
	tt := []struct {
		name    string
		opts    []ValidatorOption
		options validatorOptions
	}{
		{
			name: "defaultValidatorOptions",
			options: validatorOptions{
				omitInvalidValues: true,
				factory:           factory.StandardFactory(),
			},
		},
		{
			name: "with options",
			opts: []ValidatorOption{
				ValidatorWithPreserveInvalidValues(),
				ValidatorWithFactory(fac),
			},
			options: validatorOptions{
				omitInvalidValues: false,
				factory:           fac,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mv := NewMessageValidator(tc.opts...).(*messageValidator)
			if diff := cmp.Diff(mv.options, tc.options,
				cmp.AllowUnexported(validatorOptions{}),
				cmp.Transformer("factory", func(factory Factory) any {
					return reflect.ValueOf(factory).Pointer()
				}),
			); diff != "" {
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
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(fieldnum.RecordHeartRate),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint8),
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
							BaseType:           basetype.Uint8,
							Value:              proto.Uint8(60),
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
						field.Value = proto.Uint32(1000)
						return field
					}(),
				),
			},
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
		{
			name: "invalid utf-8 string",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("\xbd"),
				),
			},
			errs: []error{ErrInvalidUTF8String},
		},
		{
			name: "invalid utf-8 []string",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("valid utf-8 string"),
				),
				factory.CreateMesg(mesgnum.SegmentFile).WithFields(
					factory.CreateField(mesgnum.SegmentFile, fieldnum.SegmentFileLeaderActivityIdString).WithValue([]string{"valid utf-8", "\xbd"}), // valid and invalid string in array
				),
			},
			errs: []error{nil, ErrInvalidUTF8String},
		},
		{
			name: "n fields exceed allowed",
			mesgs: []proto.Message{
				{
					Num: 20,
					Fields: func() []proto.Field {
						fields := make([]proto.Field, 256)
						for i := range fields {
							fields[i] = factory.CreateField(mesgnum.Record, 255)
							fields[i].FieldBase = &proto.FieldBase{
								Num:      255,
								Name:     factory.NameUnknown,
								Type:     profile.Byte,
								BaseType: profile.Byte.BaseType(),
								Scale:    1,
								Offset:   0,
							}
							fields[i].Value = proto.Uint8(1)
						}
						return fields
					}(),
				},
			},
			errs: []error{ErrExceedMaxAllowed},
		},
		{
			name: "n developer fields exceed allowed",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "Heart Rate",
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(fieldnum.RecordHeartRate),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint8),
				}),
				{
					Num: mesgnum.Record,
					Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
					},
					DeveloperFields: func() []proto.DeveloperField {
						devFields := make([]proto.DeveloperField, 256)
						for i := range devFields {
							// dummy data
							devFields[i].DeveloperDataIndex = 0
							devFields[i].Num = 0
							devFields[i].Size = 1
							devFields[i].Name = "Heart Rate"
							devFields[i].NativeMesgNum = mesgnum.Record
							devFields[i].NativeFieldNum = fieldnum.RecordHeartRate
							devFields[i].BaseType = basetype.Uint8
							devFields[i].Value = proto.Uint8(60)
						}
						return devFields
					}(),
				},
			},
			errs: []error{nil, nil, ErrExceedMaxAllowed},
		},
		{
			name: "field value size exceed max allowed",
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue(strings.Repeat("a", 256)),
				),
			},
			errs: []error{ErrExceedMaxAllowed},
		},
		{
			name: "developer field value size exceed max allowed",
			mesgValidator: func() MessageValidator {
				mesgValidator := NewMessageValidator().(*messageValidator)
				developerDataId := proto.Message{
					Num: mesgnum.DeveloperDataId,
					Fields: []proto.Field{
						factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
					},
				}
				mesgValidator.developerDataIds = []*mesgdef.DeveloperDataId{
					mesgdef.NewDeveloperDataId(&developerDataId),
				}

				fieldDescription := proto.Message{
					Num: mesgnum.FieldDescription,
					Fields: []proto.Field{
						factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
						factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(1)),
					},
				}
				mesgValidator.fieldDescriptions = []*mesgdef.FieldDescription{
					mesgdef.NewFieldDescription(&fieldDescription),
				}
				return mesgValidator
			}(),
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.FileId).WithDeveloperFields(
					proto.DeveloperField{
						DeveloperDataIndex: 0,
						Num:                1,
						BaseType:           basetype.String,
						Value:              proto.String(strings.Repeat("a", 256)),
					},
				),
			},
			errs: []error{ErrExceedMaxAllowed},
		},
		{
			name: "valid message with developer fields invalid value",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "Heart Rate",
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(fieldnum.RecordHeartRate),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint8),
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
							BaseType:           basetype.Uint8,
							Value:              proto.Uint8(basetype.Uint8Invalid),
						},
					},
				},
			},
		},
		{
			name: "mesg contain developer field with native value scaled",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "Altitude",
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(fieldnum.RecordAltitude),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint16),
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
							NativeFieldNum:     fieldnum.RecordAltitude,
							BaseType:           basetype.Uint16,
							Value:              proto.Float64(6960.8),
						},
					},
				},
			},
		},
		{
			name: "mesg contain developer field with unknown native",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "??",
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(255),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint16),
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
							Name:               "??",
							NativeMesgNum:      mesgnum.Record,
							NativeFieldNum:     255,
							BaseType:           basetype.Uint16,
							Value:              proto.Float64(0), // Scaled value + targeted Native Field not found
						},
					},
				},
			},
			errs: []error{nil, nil, ErrValueTypeMismatch},
		},
		{
			// in Profile.xlsx (v21.133) there is no field with scale 1 and offset other than 0, but just in case in future update.
			name: "message having field with scale: 1, offset: 20",
			mesgs: []proto.Message{
				{
					Fields: []proto.Field{
						{
							FieldBase: &proto.FieldBase{
								Num:      254,
								Name:     "Unknown",
								Scale:    1,
								Offset:   20,
								Type:     profile.Uint16,
								BaseType: basetype.Uint16,
							},
							Value: proto.Float64(200),
						},
					},
				},
			},
		},
		{
			name: "valid message with developer fields has invalid value",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
					fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
					fieldnum.DeveloperDataIdApplicationId:      []byte{0, 1, 2, 3},
				}),
				factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
					fieldnum.FieldDescriptionDeveloperDataIndex:    uint8(0),
					fieldnum.FieldDescriptionFieldDefinitionNumber: uint8(0),
					fieldnum.FieldDescriptionFieldName:             "Heart Rate",
					fieldnum.FieldDescriptionNativeMesgNum:         uint16(mesgnum.Record),
					fieldnum.FieldDescriptionNativeFieldNum:        uint8(fieldnum.RecordHeartRate),
					fieldnum.FieldDescriptionFitBaseTypeId:         uint8(basetype.Uint8),
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
							BaseType:           basetype.Uint8,
							Value:              proto.Value{}, // invalid value
						},
						{
							DeveloperDataIndex: 0,
							Num:                0,
							Size:               1,
							Name:               "Heart Rate",
							NativeMesgNum:      mesgnum.Record,
							NativeFieldNum:     fieldnum.RecordHeartRate,
							BaseType:           basetype.Uint8,
							Value:              proto.Uint8(60),
						},
					},
				},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			mesgValidator := tc.mesgValidator
			if mesgValidator == nil {
				mesgValidator = NewMessageValidator()
			}

			if tc.errs == nil {
				tc.errs = make([]error, len(tc.mesgs))
			}

			for j, mesg := range tc.mesgs {
				err := mesgValidator.Validate(&mesg)
				if !errors.Is(err, tc.errs[j]) {
					t.Fatalf("expected err: %v, got: %v", tc.errs[j], err)
				}
				if err != nil {
					continue
				}
			}
		})
	}
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
			field.Value = proto.Uint32(1000)
			return field
		}(),
		func() proto.Field {
			field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
			field.IsExpandedField = true
			field.Value = proto.Uint32(1000)
			return field
		}(),
		func() proto.Field {
			field := factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedSpeed)
			field.IsExpandedField = true
			field.Value = proto.Uint32(1000)
			return field
		}(),
	)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = mesgValidator.Validate(&mesg)
	}
}
