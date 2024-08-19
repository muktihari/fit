// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// ThreeDSensorCalibration is a ThreeDSensorCalibration message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ThreeDSensorCalibration struct {
	Timestamp          time.Time          // Units: s; Whole second part of the timestamp
	OffsetCal          [3]int32           // Array: [3]; Internal calibration factors, one for each: xy, yx, zx
	OrientationMatrix  [9]int32           // Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)
	CalibrationFactor  uint32             // Calibration factor used to convert from raw ADC value to degrees, g, etc.
	CalibrationDivisor uint32             // Units: counts; Calibration factor divisor
	LevelShift         uint32             // Level shift value used to shift the ADC value back into range
	SensorType         typedef.SensorType // Indicates which sensor the calibration is for

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewThreeDSensorCalibration creates new ThreeDSensorCalibration struct based on given mesg.
// If mesg is nil, it will return ThreeDSensorCalibration with all fields being set to its corresponding invalid value.
func NewThreeDSensorCalibration(mesg *proto.Message) *ThreeDSensorCalibration {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &ThreeDSensorCalibration{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		SensorType:         typedef.SensorType(vals[0].Uint8()),
		CalibrationFactor:  vals[1].Uint32(),
		CalibrationDivisor: vals[2].Uint32(),
		LevelShift:         vals[3].Uint32(),
		OffsetCal: func() (arr [3]int32) {
			arr = [3]int32{
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
			}
			copy(arr[:], vals[4].SliceInt32())
			return arr
		}(),
		OrientationMatrix: func() (arr [9]int32) {
			arr = [9]int32{
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
				basetype.Sint32Invalid,
			}
			copy(arr[:], vals[5].SliceInt32())
			return arr
		}(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ThreeDSensorCalibration into proto.Message. If options is nil, default options will be used.
func (m *ThreeDSensorCalibration) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumThreeDSensorCalibration}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if byte(m.SensorType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.SensorType))
		fields = append(fields, field)
	}
	if m.CalibrationFactor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.CalibrationFactor)
		fields = append(fields, field)
	}
	if m.CalibrationDivisor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.CalibrationDivisor)
		fields = append(fields, field)
	}
	if m.LevelShift != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.LevelShift)
		fields = append(fields, field)
	}
	if m.OffsetCal != [3]int32{
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
	} {
		field := fac.CreateField(mesg.Num, 4)
		copied := m.OffsetCal
		field.Value = proto.SliceInt32(copied[:])
		fields = append(fields, field)
	}
	if m.OrientationMatrix != [9]int32{
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
	} {
		field := fac.CreateField(mesg.Num, 5)
		copied := m.OrientationMatrix
		field.Value = proto.SliceInt32(copied[:])
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetCalibrationFactor returns Dynamic Field interpretation of CalibrationFactor. Otherwise, returns the original value of CalibrationFactor.
//
// Based on m.SensorType:
//   - name: "accel_cal_factor", units: "g" , value: uint32(m.CalibrationFactor)
//   - name: "gyro_cal_factor", units: "deg/s" , value: uint32(m.CalibrationFactor)
//
// Otherwise:
//   - name: "calibration_factor", value: m.CalibrationFactor
func (m *ThreeDSensorCalibration) GetCalibrationFactor() (name string, value any) {
	switch m.SensorType {
	case typedef.SensorTypeAccelerometer:
		return "accel_cal_factor", uint32(m.CalibrationFactor)
	case typedef.SensorTypeGyroscope:
		return "gyro_cal_factor", uint32(m.CalibrationFactor)
	}
	return "calibration_factor", m.CalibrationFactor
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *ThreeDSensorCalibration) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// OrientationMatrixScaled return OrientationMatrix in its scaled value.
// If OrientationMatrix value is invalid, nil will be returned.
//
// Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)
func (m *ThreeDSensorCalibration) OrientationMatrixScaled() [9]float64 {
	if m.OrientationMatrix == [9]int32{
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
	} {
		return [9]float64{
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
		}
	}
	var vals [9]float64
	for i := range m.OrientationMatrix {
		if m.OrientationMatrix[i] == basetype.Sint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.OrientationMatrix[i])/65535 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of the timestamp
func (m *ThreeDSensorCalibration) SetTimestamp(v time.Time) *ThreeDSensorCalibration {
	m.Timestamp = v
	return m
}

// SetSensorType sets SensorType value.
//
// Indicates which sensor the calibration is for
func (m *ThreeDSensorCalibration) SetSensorType(v typedef.SensorType) *ThreeDSensorCalibration {
	m.SensorType = v
	return m
}

// SetCalibrationFactor sets CalibrationFactor value.
//
// Calibration factor used to convert from raw ADC value to degrees, g, etc.
func (m *ThreeDSensorCalibration) SetCalibrationFactor(v uint32) *ThreeDSensorCalibration {
	m.CalibrationFactor = v
	return m
}

// SetCalibrationDivisor sets CalibrationDivisor value.
//
// Units: counts; Calibration factor divisor
func (m *ThreeDSensorCalibration) SetCalibrationDivisor(v uint32) *ThreeDSensorCalibration {
	m.CalibrationDivisor = v
	return m
}

// SetLevelShift sets LevelShift value.
//
// Level shift value used to shift the ADC value back into range
func (m *ThreeDSensorCalibration) SetLevelShift(v uint32) *ThreeDSensorCalibration {
	m.LevelShift = v
	return m
}

// SetOffsetCal sets OffsetCal value.
//
// Array: [3]; Internal calibration factors, one for each: xy, yx, zx
func (m *ThreeDSensorCalibration) SetOffsetCal(v [3]int32) *ThreeDSensorCalibration {
	m.OffsetCal = v
	return m
}

// SetOrientationMatrix sets OrientationMatrix value.
//
// Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)
func (m *ThreeDSensorCalibration) SetOrientationMatrix(v [9]int32) *ThreeDSensorCalibration {
	m.OrientationMatrix = v
	return m
}

// SetOrientationMatrixScaled is similar to SetOrientationMatrix except it accepts a scaled value.
// This method automatically converts the given value to its [9]int32 form, discarding any applied scale and offset.
//
// Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)
func (m *ThreeDSensorCalibration) SetOrientationMatrixScaled(vs [9]float64) *ThreeDSensorCalibration {
	m.OrientationMatrix = [9]int32{
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
		basetype.Sint32Invalid,
	}
	for i := range vs {
		unscaled := (vs[i] + 0) * 65535
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint32Invalid) {
			continue
		}
		m.OrientationMatrix[i] = int32(unscaled)
	}
	return m
}

// SetDeveloperFields ThreeDSensorCalibration's DeveloperFields.
func (m *ThreeDSensorCalibration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ThreeDSensorCalibration {
	m.DeveloperFields = developerFields
	return m
}
