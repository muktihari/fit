package proto

import (
	"fmt"
	"testing"
)

func TestIsValueEqualTo(t *testing.T) {
	tt := []struct {
		field Field
		value int64
		eq    bool
	}{
		{
			field: Field{Value: uint8(89)},
			value: 89,
			eq:    true,
		},
		{
			field: Field{Value: string("fit")},
			value: 89,
			eq:    false,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v, %t", tc.value, tc.eq), func(t *testing.T) {
			if eq := tc.field.isValueEqualTo(tc.value); eq != tc.eq {
				t.Fatalf("expected: %t, got: %t", tc.eq, eq)
			}
		})
	}
}
