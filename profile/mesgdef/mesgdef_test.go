package mesgdef

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/proto"
)

func TestFieldPool(t *testing.T) {
	fields, ok := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fields)
	if !ok {
		t.Fatalf("expected ok, got not ok")
	}
}

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
