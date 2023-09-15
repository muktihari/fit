// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv

import (
	"math"
	"reflect"

	"github.com/muktihari/fit/profile/basetype"
)

func ToEnum[T ~byte](val any) T {
	return ToUint8[T](val)
}

func ToSint8[T ~int8](val any) T {
	if val == nil {
		return T(basetype.Sint8Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case int8:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Int8 {
		return T(rv.Int())
	}

	return T(basetype.Sint8Invalid)
}

func ToUint8[T ~uint8](val any) T {
	if val == nil {
		return T(basetype.Uint8Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case byte:
		return T(v)
	case bool:
		if v {
			return T(1)
		}
		return T(0)
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Uint8:
		return T(rv.Uint())
	case reflect.Bool:
		if rv.Bool() {
			return T(1)
		}
		return T(0)
	}

	return T(basetype.Uint8Invalid)
}

func ToSint16[T ~int16](val any) T {
	if val == nil {
		return T(basetype.Sint16Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case int16:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Int16 {
		return T(rv.Int())
	}

	return T(basetype.Sint16Invalid)
}

func ToUint16[T ~uint16](val any) T {
	if val == nil {
		return T(basetype.Uint16Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case uint16:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Uint16 {
		return T(rv.Uint())
	}

	return T(basetype.Uint16Invalid)
}

func ToSint32[T ~int32](val any) T {
	if val == nil {
		return T(basetype.Sint32Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case int32:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Int32 {
		return T(rv.Int())
	}

	return T(basetype.Sint32Invalid)
}

func ToUint32[T ~uint32](val any) T {
	if val == nil {
		return T(basetype.Uint32Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case uint32:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Uint32 {
		return T(rv.Uint())
	}

	return T(basetype.Uint32Invalid)
}

func ToString[T ~string](val any) T {
	if val == nil {
		return T("")
	}

	switch v := val.(type) {
	case T:
		return v
	case string:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.String {
		return T(rv.String())
	}

	return T("")
}

func ToFloat32[T ~float32](val any) T {
	if val == nil {
		return T(math.Float32frombits(basetype.Float32Invalid))
	}

	switch v := val.(type) {
	case uint32:
		return T(math.Float32frombits(v))
	case float32:
		return T(v)
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Uint32:
		return T(math.Float32frombits(uint32(rv.Uint())))
	case reflect.Float32:
		return T(rv.Float())
	}

	return T(math.Float32frombits(basetype.Float32Invalid))
}

func ToFloat64[T ~float64](val any) T {
	if val == nil {
		return T(math.Float64frombits(basetype.Float64Invalid))
	}

	switch v := val.(type) {
	case uint64:
		return T(math.Float64frombits(v))
	case float64:
		return T(v)
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Uint64:
		return T(math.Float64frombits(uint64(rv.Uint())))
	case reflect.Float64:
		return T(rv.Float())
	}

	return T(math.Float64frombits(basetype.Float64Invalid))
}

func ToUint8z[T ~uint8](val any) T {
	v := ToUint8[T](val)
	if v == T(basetype.Uint8Invalid) {
		return T(basetype.Uint8zInvalid)
	}
	return v
}

func ToUint16z[T ~uint16](val any) T {
	v := ToUint16[T](val)
	if v == T(basetype.Uint16Invalid) {
		return T(basetype.Uint16zInvalid)
	}
	return v
}

func ToUint32z[T ~uint32](val any) T {
	v := ToUint32[T](val)
	if v == T(basetype.Uint32Invalid) {
		return T(basetype.Uint32zInvalid)
	}
	return v
}

func ToByte[T ~byte](val any) T {
	return ToUint8[T](val)
}

func ToSint64[T ~int64](val any) T {
	if val == nil {
		return T(basetype.Sint64Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case int64:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Int64 {
		return T(rv.Int())
	}

	return T(basetype.Sint64Invalid)
}

func ToUint64[T ~uint64](val any) T {
	if val == nil {
		return T(basetype.Uint64Invalid)
	}

	switch v := val.(type) {
	case T:
		return v
	case uint64:
		return T(v)
	}

	if rv := reflect.ValueOf(val); rv.Kind() == reflect.Uint64 {
		return T(rv.Uint())
	}

	return T(basetype.Uint64Invalid)
}

func ToUint64z[T ~uint64](val any) T {
	v := ToUint64[T](val)
	if v == T(basetype.Uint64Invalid) {
		return T(basetype.Uint64zInvalid)
	}
	return v
}

func ToSliceEnum[T ~byte](val any) []T {
	return ToSliceUint8[T](val)
}

func ToSliceSint8[T ~int8](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []int8:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Int8 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Int()))
	}

	return vs
}

func ToSliceUint8[T ~uint8](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []uint8:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	case []bool:
		vs := make([]T, 0, len(v))
		for i := range v {
			if v[i] {
				vs = append(vs, T(1))
			} else {
				vs = append(vs, T(0))
			}
		}
		return vs

	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		switch rv.Index(i).Kind() {
		case reflect.Uint8:
			vs = append(vs, T(rv.Index(i).Uint()))
		case reflect.Bool:
			if rv.Index(i).Bool() {
				vs = append(vs, T(1))
			} else {
				vs = append(vs, T(0))
			}
		default:
			return nil
		}
	}

	return vs
}

func ToSliceSint16[T ~int16](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []int16:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Int16 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Int()))
	}

	return vs
}

func ToSliceUint16[T ~uint16](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []uint16:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Uint16 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Uint()))
	}

	return vs
}

func ToSliceSint32[T ~int32](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []int32:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Int32 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Int()))
	}

	return vs
}

func ToSliceUint32[T ~uint32](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []uint32:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Uint32 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Uint()))
	}

	return vs
}

func ToSliceFloat32[T ~float32](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []float32:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	case []uint32:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(math.Float32frombits(v[i])))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		switch rv.Index(i).Kind() {
		case reflect.Float32:
			vs = append(vs, T(rv.Index(i).Float()))
		case reflect.Uint32:
			vs = append(vs, T(math.Float32frombits(uint32(rv.Index(i).Uint()))))
		default:
			return nil
		}
	}
	return vs
}

func ToSliceFloat64[T ~float64](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []float64:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	case []uint64:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(math.Float64frombits(v[i])))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		switch rv.Index(i).Kind() {
		case reflect.Float64:
			vs = append(vs, T(rv.Index(i).Float()))
		case reflect.Uint64:
			vs = append(vs, T(math.Float64frombits(rv.Index(i).Uint())))
		default:
			return nil
		}
	}
	return vs
}

func ToSliceUint8z[T ~uint8](val any) []T {
	return ToSliceUint8[T](val)
}

func ToSliceUint16z[T ~uint16](val any) []T {
	return ToSliceUint16[T](val)
}

func ToSliceUint32z[T ~uint32](val any) []T {
	return ToSliceUint32[T](val)
}

func ToSliceByte[T ~byte](val any) []T {
	return ToSliceUint8[T](val)
}

func ToSliceSint64[T ~int64](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []int64:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Int64 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Int()))
	}

	return vs
}

func ToSliceUint64[T ~uint64](val any) []T {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case []T:
		return v
	case []uint64:
		vs := make([]T, 0, len(v))
		for i := range v {
			vs = append(vs, T(v[i]))
		}
		return vs
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	vs := make([]T, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		if rv.Index(i).Kind() != reflect.Uint64 {
			return nil
		}
		vs = append(vs, T(rv.Index(i).Uint()))
	}

	return vs
}

func ToSliceUint64z[T ~uint64](val any) []T {
	return ToSliceUint64[T](val)
}
