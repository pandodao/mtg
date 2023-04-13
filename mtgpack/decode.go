package mtgpack

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Decoder struct {
	io.Reader
}

func NewDecoder(b []byte) *Decoder {
	return &Decoder{Reader: bytes.NewReader(b)}
}

func (d *Decoder) read(b []byte) error {
	_, err := d.Read(b)
	return err
}

func (d *Decoder) ReadN(n int) ([]byte, error) {
	b := make([]byte, n)
	err := d.read(b)
	return b, err
}

func (d *Decoder) uint8() (uint8, error) {
	b, err := d.ReadN(1)
	return b[0], err
}

func (d *Decoder) uint16() (uint16, error) {
	b, err := d.ReadN(2)
	return binary.BigEndian.Uint16(b), err
}

func (d *Decoder) uint32() (uint32, error) {
	b, err := d.ReadN(4)
	return binary.BigEndian.Uint32(b), err
}

func (d *Decoder) uint64() (uint64, error) {
	b, err := d.ReadN(8)
	return binary.BigEndian.Uint64(b), err
}

func (d *Decoder) readLen() (int, error) {
	l, err := d.uint8()
	return int(l), err
}

func (d *Decoder) DecodeUint8() (uint8, error) {
	return d.uint8()
}

func (d *Decoder) DecodeUint16() (uint16, error) {
	return d.uint16()
}

func (d *Decoder) DecodeUint32() (uint32, error) {
	return d.uint32()
}

func (d *Decoder) DecodeUint64() (uint64, error) {
	return d.uint64()
}

func (d *Decoder) DecodeInt8() (int8, error) {
	u, err := d.uint8()
	return int8(u), err
}

func (d *Decoder) DecodeInt16() (int16, error) {
	u, err := d.uint16()
	return int16(u), err
}

func (d *Decoder) DecodeInt32() (int32, error) {
	u, err := d.uint32()
	return int32(u), err
}

func (d *Decoder) DecodeInt64() (int64, error) {
	u, err := d.uint64()
	return int64(u), err
}

func (d *Decoder) DecodeBool() (bool, error) {
	u, err := d.uint8()
	return u > 0, err
}

func (d *Decoder) DecodeBytes() ([]byte, error) {
	l, err := d.readLen()
	if err != nil {
		return nil, err
	}

	return d.ReadN(l)
}

func (d *Decoder) DecodeString() (string, error) {
	b, err := d.DecodeBytes()
	return bytesToString(b), err
}

func (d *Decoder) DecodeUUID() (uuid.UUID, error) {
	var id uuid.UUID

	b, err := d.ReadN(len(id))
	if err != nil {
		return uuid.Nil, err
	}

	copy(id[:], b)
	return id, nil
}

func (d *Decoder) DecodeDecimal() (decimal.Decimal, error) {
	x, err := d.DecodeInt64()
	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromInt(x).Shift(-fixedDecimalPrecision), nil
}
