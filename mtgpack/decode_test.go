package mtgpack

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecodeUUID(t *testing.T) {
	var b = make([]byte, 32)

	// typ := reflect.TypeOf(b)
	val := reflect.ValueOf(&b)

	if val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	t.Logf("is array: %t", val.Kind() == reflect.Array)

	if val.Kind() == reflect.Array {
		t.Log("elem type is", val.Type().Elem())
	}

	t.Logf("len: %d", val.Len())
}

func TestDecodeDecimalFromBytes(t *testing.T) {
	cases := []struct {
		b   []byte
		exp decimal.Decimal
	}{
		{
			b:   []byte{1, 1},
			exp: decimal.NewFromFloat(0.00000001),
		},
		{
			b:   []byte{8, 0, 0, 0, 2, 84, 11, 228, 0},
			exp: decimal.NewFromInt(100),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			dec := NewDecoder(c.b)
			d, err := dec.DecodeDecimalFromBytes()
			if err != nil {
				t.Fatal(err)
			}

			if !d.Equals(c.exp) {
				t.Fatalf("expected %v, got %v", c.exp, d)
			}
		})
	}
}
