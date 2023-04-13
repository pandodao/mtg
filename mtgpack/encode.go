package mtgpack

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	fixedDecimalPrecision int32 = 8
)

type Encoder struct {
	buf *bytes.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{buf: &bytes.Buffer{}}
}

func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

func (e *Encoder) Reset() {
	e.buf.Reset()
}

func (e *Encoder) Len() int {
	return e.buf.Len()
}

func (e *Encoder) Write(b []byte) (int, error) {
	return e.buf.Write(b)
}

func (e *Encoder) write(b []byte) error {
	_, err := e.buf.Write(b)
	return err
}

func (e *Encoder) write1(x uint8) error {
	return e.write([]byte{x})
}

func (e *Encoder) write2(x uint16) error {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], x)
	return e.write(b[:])
}

func (e *Encoder) write4(x uint32) error {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], x)
	return e.write(b[:])
}

func (e *Encoder) write8(x uint64) error {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], x)
	return e.write(b[:])
}

func (e *Encoder) writeLen(l int) error {
	if l > math.MaxUint8 {
		panic("too long")
	}

	return e.write1(uint8(l))
}

func (e *Encoder) EncodeInt(x int) error {
	return e.write4(uint32(x))
}

func (e *Encoder) EncodeInt8(x int8) error {
	return e.write1(uint8(x))
}

func (e *Encoder) EncodeInt16(x int16) error {
	return e.write2(uint16(x))
}

func (e *Encoder) EncodeInt32(x int32) error {
	return e.write4(uint32(x))
}

func (e *Encoder) EncodeInt64(x int64) error {
	return e.write8(uint64(x))
}

func (e *Encoder) EncodeUint8(x uint8) error {
	return e.write1(x)
}

func (e *Encoder) EncodeUint16(x uint16) error {
	return e.write2(x)
}

func (e *Encoder) EncodeUint32(x uint32) error {
	return e.write4(x)
}

func (e *Encoder) EncodeUint64(x uint64) error {
	return e.write8(x)
}

func (e *Encoder) EncodeUUID(x uuid.UUID) error {
	return e.write(x[:])
}

func (e *Encoder) EncodeBytes(b []byte) error {
	if err := e.writeLen(len(b)); err != nil {
		return err
	}

	return e.write(b)
}

func (e *Encoder) EncodeString(s string) error {
	return e.EncodeBytes(stringToBytes(s))
}

func (e *Encoder) EncodeDecimal(d decimal.Decimal) error {
	x := d.Shift(fixedDecimalPrecision).IntPart()
	return e.EncodeInt64(x)
}

func (e *Encoder) EncodeBool(b bool) error {
	if b {
		return e.write1(1)
	}

	return e.write1(0)
}
