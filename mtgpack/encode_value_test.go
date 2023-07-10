package mtgpack

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeValue(t *testing.T) {
	var (
		enc = NewEncoder()
		dec = &Decoder{Reader: enc.buf}
	)

	fn := func(t *testing.T, x interface{}, size int) {
		require.NoErrorf(t, EncodeValue(enc, x), "encode %T", x)
		assert.Lenf(t, enc.buf.Bytes(), size, "encode %T", x)
		y := reflect.New(reflect.TypeOf(x)).Interface()
		require.NoErrorf(t, DecodeValue(dec, y), "decode %T", x)
		assert.Equalf(t, x, reflect.ValueOf(y).Elem().Interface(), "decode %T", x)
	}

	fn(t, int8(rand.Intn(math.MaxInt8)), 1)
	fn(t, int16(rand.Intn(math.MaxInt16)), 2)
	fn(t, int32(rand.Intn(math.MaxInt32)), 4)
	fn(t, rand.Int63n(math.MaxInt64), 8)
	fn(t, uint8(rand.Intn(math.MaxInt8)), 1)
	fn(t, uint16(rand.Intn(math.MaxInt16)), 2)
	fn(t, uint32(rand.Intn(math.MaxInt32)), 4)
	fn(t, uint64(rand.Int63n(math.MaxInt64)), 8)
	fn(t, time.Date(2021, 45, 56, 34, 76, 92, 66565, time.Local), 8)
	fn(t, "foo", 4)
	fn(t, rand.Intn(2) == 1, 1)
	fn(t, uuid.New(), 16)
	fn(t, decimal.NewFromInt(1345).Shift(-8), 8)

	assert.Emptyf(t, enc.buf.Len(), "encoder has not remaining bytes")
}
