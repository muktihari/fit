package filedef

import (
	"reflect"
	"testing"

	"github.com/muktihari/fit/profile/typedef"
)

func TestWithFileFunc(t *testing.T) {
	l := NewListener(
		WithFileFunc(typedef.FileActivity, func() File { return nil }),
	)

	// Test whether readOnlyFileSets is read only and should not be altered when assigning custom File.
	// WithFileFunc should make a clone of it.
	if reflect.ValueOf(readOnlyFileSets).Pointer() == reflect.ValueOf(l.options.fileSets).Pointer() {
		t.Fatalf("listener's fileSets still referenced to readOnlyFileSets")
	}

	// Ensure listener's fileSets for FileActivity is being replaced with function above.
	if file := l.options.fileSets[typedef.FileActivity](); file != nil {
		t.Fatalf("expected nil, got: %v", file)
	}

	// Ensure readOnlyFileSets still returning Activity, and not being altered.
	file := readOnlyFileSets[typedef.FileActivity]()
	if _, ok := file.(*Activity); !ok {
		t.Fatalf("expected *Activity, got %T", file)
	}
}
