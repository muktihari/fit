package kit_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit"
)

func TestPtr(t *testing.T) {
	val := float64(10)
	ptr := kit.Ptr(val)

	if diff := cmp.Diff(val, *ptr); diff != "" {
		t.Fatal(diff)
	}
}
