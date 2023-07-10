package mtgpack

import (
	"bytes"
	"encoding/binary"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Decoder struct {
	io.Reader
}

// NewDecoder returns a new Decoder with the provided byte slice as its input.
func NewDecoder(b []byte) *Decoder {
	return &Decoder{Reader: bytes.NewReader(b)}
}

// read reads from the underlying input and fills the provided byte slice.
func (d *Decoder) read(b []byte) error {
	_, err := d.Read(b)
	return err
}

// ReadN reads n bytes from the underlying input.
func (d *Decoder) ReadN(n int) ([]byte, error) {
	b := make([]byte, n)
	err := d.read(b)
	return b, err
}

// uint8 reads a uint8 from the underlying input.
func (d *Decoder) uint8() (uint8, error) {
	b, err := d.ReadN(1)
	return b[0], err
}

// uint16 reads a uint16 from the underlying input.
func (d *Decoder) uint16() (uint16, error) {
	b, err := d.ReadN(2)
	return binary.BigEndian.Uint16(b), err
}

// uint32 reads a uint32 from the underlying input.
func (d *Decoder) uint32() (uint32, error) {
	b, err := d.ReadN(4)
	return binary.BigEndian.Uint32(b), err
}

// uint64 reads a uint64 from the underlying input.
func (d *Decoder) uint64() (uint64, error) {
	b, err := d.ReadN(8)
	return binary.BigEndian.Uint64(b), err
}

// readLen reads the length of a byte array from the underlying input.
func (d *Decoder) readLen() (int, error) {
	l, err := d.uint8()
	return int(l), err
}

// DecodeUint8 decodes a uint8 from the input.
func (d *Decoder) DecodeUint8() (uint8, error) {
	return d.uint8()
}

// DecodeUint16 decodes a uint16 from the input.
func (d *Decoder) DecodeUint16() (uint16, error) {
	return d.uint16()
}

// DecodeUint32 decodes a uint32 from the input.
func (d *Decoder) DecodeUint32() (uint32, error) {
	return d.uint32()
}

// DecodeUint64 decodes a uint64 from the input.
func (d *Decoder) DecodeUint64() (uint64, error) {
	return d.uint64()
}

// DecodeInt8 decodes an int8 from the input.
func (d *Decoder) DecodeInt8() (int8, error) {
	u, err := d.uint8()
	return int8(u), err
}

// DecodeInt16 decodes an int16 from the input.
func (d *Decoder) DecodeInt16() (int16, error) {
	u, err := d.uint16()
	return int16(u), err
}

// DecodeInt32 decodes an int32 from the input.
func (d *Decoder) DecodeInt32() (int32, error) {
	u, err := d.uint32()
	return int32(u), err
}

// DecodeInt64 decodes an int64 from the input.
func (d *Decoder) DecodeInt64() (int64, error) {
	u, err := d.uint64()
	return int64(u), err
}

// DecodeBool decodes a bool from the input.
func (d *Decoder) DecodeBool() (bool, error) {
	u, err := d.uint8()
	return u > 0, err
}

// DecodeBytes decodes a byte array from the input.
func (d *Decoder) DecodeBytes() ([]byte, error) {
	l, err := d.readLen()
	if err != nil {
		return nil, err
	}

	return d.ReadN(l)
}

// DecodeString decodes a string from the input.
func (d *Decoder) DecodeString() (string, error) {
	b, err := d.DecodeBytes()
	return bytesToString(b), err
}

// DecodeUUID decodes a UUID from the input.
func (d *Decoder) DecodeUUID() (uuid.UUID, error) {
	var id uuid.UUID

	b, err := d.ReadN(len(id))
	if err != nil {
		return uuid.Nil, err
	}

	copy(id[:], b)
	return id, nil
}

// DecodeDecimal decodes a decimal.Decimal number from the input.
func (d *Decoder) DecodeDecimal() (decimal.Decimal, error) {
	x, err := d.DecodeInt64()
	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromInt(x).Shift(-fixedDecimalPrecision), nil
}

// DecodeTime decodes a time.Time from the input.
func (d *Decoder) DecodeTime() (time.Time, error) {
	x, err := d.DecodeInt64()
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, x), nil
}
