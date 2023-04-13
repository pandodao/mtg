package mtgpack

import (
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
)

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

	if typ == decimalType {
		d := val.Interface().(decimal.Decimal)
		return e.EncodeDecimal(d)
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
