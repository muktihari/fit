// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// DeviceInfo is a DeviceInfo message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type DeviceInfo struct {
	Timestamp           time.Time // Units: s
	Descriptor          string    // Used to describe the sensor or location
	ProductName         string    // Optional free form string to indicate the devices name or model
	SerialNumber        uint32    // Base: uint32z
	CumOperatingTime    uint32    // Units: s; Reset by new battery or charge.
	Manufacturer        typedef.Manufacturer
	Product             uint16
	SoftwareVersion     uint16 // Scale: 100
	BatteryVoltage      uint16 // Scale: 256; Units: V
	AntDeviceNumber     uint16 // Base: uint16z
	DeviceIndex         typedef.DeviceIndex
	DeviceType          uint8
	HardwareVersion     uint8
	BatteryStatus       typedef.BatteryStatus
	SensorPosition      typedef.BodyLocation // Indicates the location of the sensor
	AntTransmissionType uint8                // Base: uint8z
	AntNetwork          typedef.AntNetwork
	SourceType          typedef.SourceType
	BatteryLevel        uint8 // Units: %

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewDeviceInfo creates new DeviceInfo struct based on given mesg.
// If mesg is nil, it will return DeviceInfo with all fields being set to its corresponding invalid value.
func NewDeviceInfo(mesg *proto.Message) *DeviceInfo {
	m := new(DeviceInfo)
	m.Reset(mesg)
	return m
}

// Reset resets all DeviceInfo's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *DeviceInfo) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = DeviceInfo{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		DeviceIndex:         typedef.DeviceIndex(vals[0].Uint8()),
		DeviceType:          vals[1].Uint8(),
		Manufacturer:        typedef.Manufacturer(vals[2].Uint16()),
		SerialNumber:        vals[3].Uint32z(),
		Product:             vals[4].Uint16(),
		SoftwareVersion:     vals[5].Uint16(),
		HardwareVersion:     vals[6].Uint8(),
		CumOperatingTime:    vals[7].Uint32(),
		BatteryVoltage:      vals[10].Uint16(),
		BatteryStatus:       typedef.BatteryStatus(vals[11].Uint8()),
		SensorPosition:      typedef.BodyLocation(vals[18].Uint8()),
		Descriptor:          vals[19].String(),
		AntTransmissionType: vals[20].Uint8z(),
		AntDeviceNumber:     vals[21].Uint16z(),
		AntNetwork:          typedef.AntNetwork(vals[22].Uint8()),
		SourceType:          typedef.SourceType(vals[25].Uint8()),
		ProductName:         vals[27].String(),
		BatteryLevel:        vals[32].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts DeviceInfo into proto.Message. If options is nil, default options will be used.
func (m *DeviceInfo) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumDeviceInfo}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.DeviceIndex != typedef.DeviceIndexInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if m.DeviceType != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.DeviceType)
		fields = append(fields, field)
	}
	if m.Manufacturer != typedef.ManufacturerInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.SerialNumber != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}
	if m.SoftwareVersion != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.SoftwareVersion)
		fields = append(fields, field)
	}
	if m.HardwareVersion != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.HardwareVersion)
		fields = append(fields, field)
	}
	if m.CumOperatingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint32(m.CumOperatingTime)
		fields = append(fields, field)
	}
	if m.BatteryVoltage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(m.BatteryVoltage)
		fields = append(fields, field)
	}
	if m.BatteryStatus != typedef.BatteryStatusInvalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(uint8(m.BatteryStatus))
		fields = append(fields, field)
	}
	if m.SensorPosition != typedef.BodyLocationInvalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Uint8(byte(m.SensorPosition))
		fields = append(fields, field)
	}
	if m.Descriptor != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.String(m.Descriptor)
		fields = append(fields, field)
	}
	if m.AntTransmissionType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.Uint8(m.AntTransmissionType)
		fields = append(fields, field)
	}
	if m.AntDeviceNumber != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint16(m.AntDeviceNumber)
		fields = append(fields, field)
	}
	if m.AntNetwork != typedef.AntNetworkInvalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint8(byte(m.AntNetwork))
		fields = append(fields, field)
	}
	if m.SourceType != typedef.SourceTypeInvalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint8(byte(m.SourceType))
		fields = append(fields, field)
	}
	if m.ProductName != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = proto.String(m.ProductName)
		fields = append(fields, field)
	}
	if m.BatteryLevel != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = proto.Uint8(m.BatteryLevel)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetDeviceType returns Dynamic Field interpretation of DeviceType. Otherwise, returns the original value of DeviceType.
//
// Based on m.SourceType:
//   - name: "ble_device_type", value: typedef.BleDeviceType(m.DeviceType)
//   - name: "antplus_device_type", value: typedef.AntplusDeviceType(m.DeviceType)
//   - name: "ant_device_type", value: uint8(m.DeviceType)
//   - name: "local_device_type", value: typedef.LocalDeviceType(m.DeviceType)
//
// Otherwise:
//   - name: "device_type", value: m.DeviceType
func (m *DeviceInfo) GetDeviceType() (name string, value any) {
	switch m.SourceType {
	case typedef.SourceTypeBluetoothLowEnergy:
		return "ble_device_type", typedef.BleDeviceType(m.DeviceType)
	case typedef.SourceTypeAntplus:
		return "antplus_device_type", typedef.AntplusDeviceType(m.DeviceType)
	case typedef.SourceTypeAnt:
		return "ant_device_type", uint8(m.DeviceType)
	case typedef.SourceTypeLocal:
		return "local_device_type", typedef.LocalDeviceType(m.DeviceType)
	}
	return "device_type", m.DeviceType
}

// GetProduct returns Dynamic Field interpretation of Product. Otherwise, returns the original value of Product.
//
// Based on m.Manufacturer:
//   - name: "favero_product", value: typedef.FaveroProduct(m.Product)
//   - name: "garmin_product", value: typedef.GarminProduct(m.Product)
//
// Otherwise:
//   - name: "product", value: m.Product
func (m *DeviceInfo) GetProduct() (name string, value any) {
	switch m.Manufacturer {
	case typedef.ManufacturerFaveroElectronics:
		return "favero_product", typedef.FaveroProduct(m.Product)
	case typedef.ManufacturerGarmin, typedef.ManufacturerDynastream, typedef.ManufacturerDynastreamOem, typedef.ManufacturerTacx:
		return "garmin_product", typedef.GarminProduct(m.Product)
	}
	return "product", m.Product
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *DeviceInfo) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SoftwareVersionScaled return SoftwareVersion in its scaled value.
// If SoftwareVersion value is invalid, float64 invalid value will be returned.
//
// Scale: 100
func (m *DeviceInfo) SoftwareVersionScaled() float64 {
	if m.SoftwareVersion == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.SoftwareVersion)/100 - 0
}

// BatteryVoltageScaled return BatteryVoltage in its scaled value.
// If BatteryVoltage value is invalid, float64 invalid value will be returned.
//
// Scale: 256; Units: V
func (m *DeviceInfo) BatteryVoltageScaled() float64 {
	if m.BatteryVoltage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BatteryVoltage)/256 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *DeviceInfo) SetTimestamp(v time.Time) *DeviceInfo {
	m.Timestamp = v
	return m
}

// SetDeviceIndex sets DeviceIndex value.
func (m *DeviceInfo) SetDeviceIndex(v typedef.DeviceIndex) *DeviceInfo {
	m.DeviceIndex = v
	return m
}

// SetDeviceType sets DeviceType value.
func (m *DeviceInfo) SetDeviceType(v uint8) *DeviceInfo {
	m.DeviceType = v
	return m
}

// SetManufacturer sets Manufacturer value.
func (m *DeviceInfo) SetManufacturer(v typedef.Manufacturer) *DeviceInfo {
	m.Manufacturer = v
	return m
}

// SetSerialNumber sets SerialNumber value.
//
// Base: uint32z
func (m *DeviceInfo) SetSerialNumber(v uint32) *DeviceInfo {
	m.SerialNumber = v
	return m
}

// SetProduct sets Product value.
func (m *DeviceInfo) SetProduct(v uint16) *DeviceInfo {
	m.Product = v
	return m
}

// SetSoftwareVersion sets SoftwareVersion value.
//
// Scale: 100
func (m *DeviceInfo) SetSoftwareVersion(v uint16) *DeviceInfo {
	m.SoftwareVersion = v
	return m
}

// SetSoftwareVersionScaled is similar to SetSoftwareVersion except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100
func (m *DeviceInfo) SetSoftwareVersionScaled(v float64) *DeviceInfo {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.SoftwareVersion = uint16(basetype.Uint16Invalid)
		return m
	}
	m.SoftwareVersion = uint16(unscaled)
	return m
}

// SetHardwareVersion sets HardwareVersion value.
func (m *DeviceInfo) SetHardwareVersion(v uint8) *DeviceInfo {
	m.HardwareVersion = v
	return m
}

// SetCumOperatingTime sets CumOperatingTime value.
//
// Units: s; Reset by new battery or charge.
func (m *DeviceInfo) SetCumOperatingTime(v uint32) *DeviceInfo {
	m.CumOperatingTime = v
	return m
}

// SetBatteryVoltage sets BatteryVoltage value.
//
// Scale: 256; Units: V
func (m *DeviceInfo) SetBatteryVoltage(v uint16) *DeviceInfo {
	m.BatteryVoltage = v
	return m
}

// SetBatteryVoltageScaled is similar to SetBatteryVoltage except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 256; Units: V
func (m *DeviceInfo) SetBatteryVoltageScaled(v float64) *DeviceInfo {
	unscaled := (v + 0) * 256
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BatteryVoltage = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BatteryVoltage = uint16(unscaled)
	return m
}

// SetBatteryStatus sets BatteryStatus value.
func (m *DeviceInfo) SetBatteryStatus(v typedef.BatteryStatus) *DeviceInfo {
	m.BatteryStatus = v
	return m
}

// SetSensorPosition sets SensorPosition value.
//
// Indicates the location of the sensor
func (m *DeviceInfo) SetSensorPosition(v typedef.BodyLocation) *DeviceInfo {
	m.SensorPosition = v
	return m
}

// SetDescriptor sets Descriptor value.
//
// Used to describe the sensor or location
func (m *DeviceInfo) SetDescriptor(v string) *DeviceInfo {
	m.Descriptor = v
	return m
}

// SetAntTransmissionType sets AntTransmissionType value.
//
// Base: uint8z
func (m *DeviceInfo) SetAntTransmissionType(v uint8) *DeviceInfo {
	m.AntTransmissionType = v
	return m
}

// SetAntDeviceNumber sets AntDeviceNumber value.
//
// Base: uint16z
func (m *DeviceInfo) SetAntDeviceNumber(v uint16) *DeviceInfo {
	m.AntDeviceNumber = v
	return m
}

// SetAntNetwork sets AntNetwork value.
func (m *DeviceInfo) SetAntNetwork(v typedef.AntNetwork) *DeviceInfo {
	m.AntNetwork = v
	return m
}

// SetSourceType sets SourceType value.
func (m *DeviceInfo) SetSourceType(v typedef.SourceType) *DeviceInfo {
	m.SourceType = v
	return m
}

// SetProductName sets ProductName value.
//
// Optional free form string to indicate the devices name or model
func (m *DeviceInfo) SetProductName(v string) *DeviceInfo {
	m.ProductName = v
	return m
}

// SetBatteryLevel sets BatteryLevel value.
//
// Units: %
func (m *DeviceInfo) SetBatteryLevel(v uint8) *DeviceInfo {
	m.BatteryLevel = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *DeviceInfo) SetUnknownFields(unknownFields ...proto.Field) *DeviceInfo {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *DeviceInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *DeviceInfo {
	m.DeveloperFields = developerFields
	return m
}
