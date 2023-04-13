package mtgpack

import (
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
	t.Run("int8", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = int8(rand.Intn(math.MaxInt8))
			y int8
		)

		require.NoErrorf(t, e.EncodeInt8(x), "encode int8")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode int8")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("int16", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = int16(rand.Intn(math.MaxInt16))
			y int16
		)

		require.NoErrorf(t, e.EncodeInt16(x), "encode int16")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode int16")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("int32", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = rand.Int31n(math.MaxInt32)
			y int32
		)

		require.NoErrorf(t, e.EncodeInt32(x), "encode int32")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode int32")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("int64", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = rand.Int63n(math.MaxInt64)
			y int64
		)

		require.NoErrorf(t, e.EncodeInt64(x), "encode int64")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode int64")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("uint8", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = uint8(rand.Intn(math.MaxUint8))
			y uint8
		)

		require.NoErrorf(t, e.EncodeUint8(x), "encode uint8")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode uint8")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("uint16", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = uint16(rand.Intn(math.MaxUint16))
			y uint16
		)

		require.NoErrorf(t, e.EncodeUint16(x), "encode uint16")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode uint16")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("uint32", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = uint32(rand.Intn(math.MaxUint32))
			y uint32
		)

		require.NoErrorf(t, e.EncodeUint32(x), "encode uint32")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode uint32")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("uint64", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = uint64(rand.Int63n(math.MaxUint32))
			y uint64
		)

		require.NoErrorf(t, e.EncodeUint64(x), "encode uint64")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode uint64")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("string", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = "hello world"
			y string
		)

		require.NoErrorf(t, e.EncodeString(x), "encode string")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode string")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("uuid", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = uuid.New()
			y uuid.UUID
		)

		require.NoErrorf(t, e.EncodeUUID(x), "encode uuid")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode uuid")
		assert.Equal(t, x, y)
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})

	t.Run("decimal", func(t *testing.T) {
		var (
			e = NewEncoder()
			x = decimal.NewFromFloat(123.456)
			y decimal.Decimal
		)

		require.NoErrorf(t, e.EncodeDecimal(x), "encode decimal")
		d := NewDecoder(e.Bytes())
		require.NoErrorf(t, DecodeValue(d, &y), "decode decimal")
		assert.Equal(t, x.String(), y.String())
		assert.Emptyf(t, d.r.Len(), "decoder has not remaining bytes")
	})
}
