package mtgpack

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

func (e *Encoder) EncodeValue(v interface{}) error {
	return EncodeValue(e, v)
}

func (e *Encoder) EncodeValues(values ...interface{}) error {
	return EncodeValues(e, values...)
}

func EncodeValues(e *Encoder, values ...interface{}) error {
	for _, v := range values {
		if err := EncodeValue(e, v); err != nil {
			return err
		}
	}

	return nil
}

// EncodeValue encode a value to the encoder.
func EncodeValue(e *Encoder, v interface{}) error {
	val := reflect.ValueOf(v)
	typ := val.Type()

	if typ.Implements(customEncoderType) {
		encoder := val.Interface().(CustomEncoder)
		return encoder.EncodeMtg(e)
	}

	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	switch typ {
	case decimalType:
		d := val.Interface().(decimal.Decimal)
		return e.EncodeDecimal(d)
	case timeType:
		t := val.Interface().(time.Time)
		return e.EncodeTime(t)
	}

	// encode uuid.UUID
	if size, ok := isByteArray(val); ok {
		return encodeByteArrayValue(e, val, size)
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int64:
		return e.EncodeInt64(val.Int())
	case reflect.Int8:
		return e.EncodeInt8(int8(val.Int()))
	case reflect.Int16:
		return e.EncodeInt16(int16(val.Int()))
	case reflect.Int32:
		return e.EncodeInt32(int32(val.Int()))
	case reflect.Uint, reflect.Uint64:
		return e.EncodeUint64(val.Uint())
	case reflect.Uint8:
		return e.EncodeUint8(uint8(val.Uint()))
	case reflect.Uint16:
		return e.EncodeUint16(uint16(val.Uint()))
	case reflect.Uint32:
		return e.EncodeUint32(uint32(val.Uint()))
	case reflect.Bool:
		return e.EncodeBool(val.Bool())
	case reflect.String:
		return e.EncodeString(val.String())
	}

	return fmt.Errorf("unsupported type: %s", typ)
}

func encodeByteArrayValue(e *Encoder, val reflect.Value, size int) error {
	b := make([]byte, size)
	reflect.Copy(reflect.ValueOf(b), val)
	_, err := e.Write(b)
	return err
}
