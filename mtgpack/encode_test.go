package mtgpack

import (
	"bytes"
	"testing"

	"github.com/shopspring/decimal"
)

func TestEncodeDecimalToBytes(t *testing.T) {
	cases := []struct {
		dec decimal.Decimal
		exp []byte
	}{
		{
			dec: decimal.NewFromFloat(0.00000001),
			exp: []byte{1, 1},
		},
		{
			dec: decimal.NewFromFloat(100),
			exp: []byte{8, 0, 0, 0, 2, 84, 11, 228, 0},
		},
	}

	for _, c := range cases {
		t.Run(c.dec.String(), func(t *testing.T) {
			enc := NewEncoder()
			err := enc.EncodeDecimalToBytes(c.dec)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(enc.Bytes(), c.exp) {
				t.Fatalf("expected %v, got %v", c.exp, enc.Bytes())
			}
		})
	}
}
