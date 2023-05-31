package mtgpack

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// fixedDecimalPrecision defines the fixed decimal precision to 8.
const (
	fixedDecimalPrecision int32 = 8
)

// Encoder provides methods for encoding different data types into a byte buffer.
type Encoder struct {
	buf *bytes.Buffer // the byte buffer where encoded data is written
}

// NewEncoder constructs and returns a new Encoder.
func NewEncoder() *Encoder {
	return &Encoder{buf: &bytes.Buffer{}}
}

// Bytes returns the current contents of the Encoder's buffer as a byte slice.
func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

// Reset resets the Encoder's buffer to an empty state.
func (e *Encoder) Reset() {
	e.buf.Reset()
}

// Len returns the number of bytes currently written to the buffer.
func (e *Encoder) Len() int {
	return e.buf.Len()
}

// Write implements io.Writer interface and writes the given bytes to the buffer.
func (e *Encoder) Write(b []byte) (int, error) {
	return e.buf.Write(b)
}

// write writes the given bytes to the buffer.
func (e *Encoder) write(b []byte) error {
	_, err := e.buf.Write(b)
	return err
}

// write1 writes the given byte to the buffer.
func (e *Encoder) write1(x uint8) error {
	return e.write([]byte{x})
}

// write2 writes the given uint16 to the buffer in big-endian format.
func (e *Encoder) write2(x uint16) error {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], x)
	return e.write(b[:])
}

// write4 writes the given uint32 to the buffer in big-endian format.
func (e *Encoder) write4(x uint32) error {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], x)
	return e.write(b[:])
}

// write8 writes the given uint64 to the buffer in big-endian format.
func (e *Encoder) write8(x uint64) error {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], x)
	return e.write(b[:])
}

// writeLen writes the given length as an uint8 to the buffer.
func (e *Encoder) writeLen(l int) error {
	if l > math.MaxUint8 {
		panic("too long")
	}

	return e.write1(uint8(l))
}

// EncodeInt encodes the given int into the buffer as a uint32.
func (e *Encoder) EncodeInt(x int) error {
	return e.write4(uint32(x))
}

// EncodeInt8 encodes the given int8 into the buffer as a uint8.
func (e *Encoder) EncodeInt8(x int8) error {
	return e.write1(uint8(x))
}

// EncodeInt16 encodes the given int16 into the buffer as a uint16.
func (e *Encoder) EncodeInt16(x int16) error {
	return e.write2(uint16(x))
}

// EncodeInt32 encodes the given int32 into the buffer as a uint32.
func (e *Encoder) EncodeInt32(x int32) error {
	return e.write4(uint32(x))
}

// EncodeInt64 encodes the given int64 into the buffer as a uint64.
func (e *Encoder) EncodeInt64(x int64) error {
	return e.write8(uint64(x))
}

// EncodeUint8 encodes the given uint8 into the buffer.
func (e *Encoder) EncodeUint8(x uint8) error {
	return e.write1(x)
}

// EncodeUint16 encodes the given uint16 into the buffer.
func (e *Encoder) EncodeUint16(x uint16) error {
	return e.write2(x)
}

// EncodeUint32 encodes the given uint32 into the buffer.
func (e *Encoder) EncodeUint32(x uint32) error {
	return e.write4(x)
}

// EncodeUint64 encodes the given uint64 into the buffer.
func (e *Encoder) EncodeUint64(x uint64) error {
	return e.write8(x)
}

// EncodeUUID encodes the given UUID into the buffer.
func (e *Encoder) EncodeUUID(x uuid.UUID) error {
	return e.write(x[:])
}

// EncodeBytes encodes the given byte slice into the buffer. It writes the length of the
// slice as a uint8 followed by the slice bytes.
func (e *Encoder) EncodeBytes(b []byte) error {
	if err := e.writeLen(len(b)); err != nil {
		return err
	}

	return e.write(b)
}

// EncodeString encodes the given string into the buffer. It first converts the string to bytes
// and then encodes the resulting byte slice using EncodeBytes.
func (e *Encoder) EncodeString(s string) error {
	return e.EncodeBytes(stringToBytes(s))
}

// EncodeDecimal encodes the given decimal into the buffer. It first shifts the decimal
// by fixedDecimalPrecision and then encodes the resulting int64 using EncodeInt64.
func (e *Encoder) EncodeDecimal(d decimal.Decimal) error {
	x := d.Shift(fixedDecimalPrecision).IntPart()
	return e.EncodeInt64(x)
}

// EncodeBool encodes the given bool into the buffer as a single byte (1 for true, 0 for false).
func (e *Encoder) EncodeBool(b bool) error {
	if b {
		return e.write1(1)
	}

	return e.write1(0)
}
