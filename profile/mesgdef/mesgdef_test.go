// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
)

func TestDefaultOptions(t *testing.T) {
	options := DefaultOptions()
	if diff := cmp.Diff(options, &Options{
		Factory:               factory.StandardFactory(),
		IncludeExpandedFields: false,
	}, cmp.Transformer("Factory", func(fac Factory) uintptr {
		return uintptr(reflect.ValueOf(fac).UnsafePointer())
	}),
	); diff != "" {
		t.Fatal(diff)
	}
}

func TestUnsafeCast(t *testing.T) {
	attitudeValidities := []typedef.AttitudeValidity{
		typedef.AttitudeValidityNoGps,
		typedef.AttitudeValidityHwFail,
		typedef.AttitudeValiditySolutionCoasting,
	}
	mesg := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
		fieldnum.AviationAttitudeValidity: attitudeValidities,
	})

	aviationAttitude := NewAviationAttitude(&mesg)
	newMesg := aviationAttitude.ToMesg(nil)

	newAttitudeValidities := newMesg.FieldValueByNum(fieldnum.AviationAttitudeValidity).SliceUint16()

	if len(attitudeValidities) != len(newAttitudeValidities) {
		t.Fatalf("expected len: %d, got: %d", len(attitudeValidities), len(newAttitudeValidities))
	}

	for i := range attitudeValidities {
		if attitudeValidities[i] != typedef.AttitudeValidity(newAttitudeValidities[i]) {
			t.Errorf("[%d] expected: %v, got: %v", i, attitudeValidities[i], newAttitudeValidities[i])
		}
	}
}
