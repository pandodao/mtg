package mtgpack

import (
	"bytes"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsByteArray(t *testing.T) {
	t.Run("bytes array", func(t *testing.T) {
		var b [8]byte
		length, ok := isByteArray(reflect.ValueOf(b))
		assert.True(t, ok)
		assert.Equal(t, len(b), length)
	})

	t.Run("uuid", func(t *testing.T) {
		var id uuid.UUID
		length, ok := isByteArray(reflect.ValueOf(id))
		assert.True(t, ok)
		assert.Equal(t, len(id), length)
	})
}

func TestDecodeValue(t *testing.T) {
	var (
		buf bytes.Buffer
		enc = &Encoder{buf: &buf}
		dec = &Decoder{Reader: &buf}
	)

	t.Run("int8", func(t *testing.T) {
		var x, y int8
		x = int8(rand.Intn(math.MaxInt8))

		require.NoErrorf(t, enc.EncodeInt8(x), "encode int8")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode int8")
		assert.Equal(t, x, y)
	})

	t.Run("int16", func(t *testing.T) {
		var x, y int16
		x = int16(rand.Intn(math.MaxInt16))

		require.NoErrorf(t, enc.EncodeInt16(x), "encode int16")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode int16")
		assert.Equal(t, x, y)
	})

	t.Run("int32", func(t *testing.T) {
		var x, y int32
		x = int32(rand.Intn(math.MaxInt32))

		require.NoErrorf(t, enc.EncodeInt32(x), "encode int32")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode int32")
		assert.Equal(t, x, y)
	})

	t.Run("int64", func(t *testing.T) {
		var x, y int64
		x = int64(rand.Intn(math.MaxInt64))

		require.NoErrorf(t, enc.EncodeInt64(x), "encode int64")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode int64")
		assert.Equal(t, x, y)
	})

	t.Run("uint8", func(t *testing.T) {
		var x, y uint8
		x = uint8(rand.Intn(math.MaxUint8))

		require.NoErrorf(t, enc.EncodeUint8(x), "encode uint8")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode uint8")
		assert.Equal(t, x, y)
	})

	t.Run("uint16", func(t *testing.T) {
		var x, y uint16
		x = uint16(rand.Intn(math.MaxUint16))

		require.NoErrorf(t, enc.EncodeUint16(x), "encode uint16")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode uint16")
		assert.Equal(t, x, y)
	})

	t.Run("uint32", func(t *testing.T) {
		var x, y uint32
		x = uint32(rand.Intn(math.MaxUint32))

		require.NoErrorf(t, enc.EncodeUint32(x), "encode uint32")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode uint32")
		assert.Equal(t, x, y)
	})

	t.Run("uint64", func(t *testing.T) {
		var x, y uint64
		x = uint64(rand.Int63n(math.MaxUint32))

		require.NoErrorf(t, enc.EncodeUint64(x), "encode uint64")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode uint64")
		assert.Equal(t, x, y)
	})

	t.Run("string", func(t *testing.T) {
		var x, y string
		x = "hello world"

		require.NoErrorf(t, enc.EncodeString(x), "encode string")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode string")
		assert.Equal(t, x, y)
	})

	t.Run("uuid", func(t *testing.T) {
		var x, y uuid.UUID
		x = uuid.New()

		require.NoErrorf(t, enc.EncodeUUID(x), "encode uuid")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode uuid")
		assert.Equal(t, x, y)
	})

	t.Run("decimal", func(t *testing.T) {
		var x, y decimal.Decimal
		x = decimal.NewFromFloat(1.234)

		require.NoErrorf(t, enc.EncodeDecimal(x), "encode decimal")
		require.NoErrorf(t, DecodeValue(dec, &y), "decode decimal")
		assert.Equal(t, x.String(), y.String())
	})

	assert.Emptyf(t, buf.Len(), "decoder has not remaining bytes")
}
