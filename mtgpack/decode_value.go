package mtgpack

import (
	"fmt"
	"reflect"
)

func (d *Decoder) DecodeValue(v interface{}) error {
	return DecodeValue(d, v)
}

func (d *Decoder) DecodeValues(values ...interface{}) error {
	return DecodeValues(d, values...)
}

func DecodeValues(d *Decoder, values ...interface{}) error {
	for _, v := range values {
		if err := DecodeValue(d, v); err != nil {
			return err
		}
	}

	return nil
}

func DecodeValue(d *Decoder, v interface{}) error {
	val := reflect.ValueOf(v)
	typ := val.Type()

	if typ.Implements(customDecoderType) {
		decoder := val.Interface().(CustomDecoder)
		return decoder.DecodeMtg(d)
	}

	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	if !val.CanSet() {
		return fmt.Errorf("cannot set value: %s", typ)
	}

	if typ == decimalType {
		return decodeDecimalValue(d, val)
	}

	// decode uuid.UUID
	if size, ok := isByteArray(val); ok {
		return decodeByteArrayValue(d, val, size)
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int64:
		return decodeInt64Value(d, val)
	case reflect.Int8:
		return decodeInt8Value(d, val)
	case reflect.Int16:
		return decodeInt16Value(d, val)
	case reflect.Int32:
		return decodeInt32Value(d, val)
	case reflect.Uint, reflect.Uint64:
		return decodeUint64Value(d, val)
	case reflect.Uint8:
		return decodeUint8Value(d, val)
	case reflect.Uint16:
		return decodeUint16Value(d, val)
	case reflect.Uint32:
		return decodeUint32Value(d, val)
	case reflect.Bool:
		return decodeBoolValue(d, val)
	case reflect.String:
		return decodeStringValue(d, val)
	}

	return fmt.Errorf("unsupported type: %s", typ)
}

func decodeInt8Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeInt8()
	if err != nil {
		return err
	}

	val.SetInt(int64(i))
	return nil
}

func decodeInt16Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeInt16()
	if err != nil {
		return err
	}

	val.SetInt(int64(i))
	return nil
}

func decodeInt32Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeInt32()
	if err != nil {
		return err
	}

	val.SetInt(int64(i))
	return nil
}

func decodeInt64Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeInt64()
	if err != nil {
		return err
	}

	val.SetInt(i)
	return nil
}

func decodeUint8Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeUint8()
	if err != nil {
		return err
	}

	val.SetUint(uint64(i))
	return nil
}

func decodeUint16Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeUint16()
	if err != nil {
		return err
	}

	val.SetUint(uint64(i))
	return nil
}

func decodeUint32Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeUint32()
	if err != nil {
		return err
	}

	val.SetUint(uint64(i))
	return nil
}

func decodeUint64Value(d *Decoder, val reflect.Value) error {
	i, err := d.DecodeUint64()
	if err != nil {
		return err
	}

	val.SetUint(i)
	return nil
}

func decodeStringValue(d *Decoder, val reflect.Value) error {
	s, err := d.DecodeString()
	if err != nil {
		return err
	}

	val.SetString(s)
	return nil
}

func decodeBoolValue(d *Decoder, val reflect.Value) error {
	b, err := d.DecodeBool()
	if err != nil {
		return err
	}

	val.SetBool(b)
	return nil
}

func decodeDecimalValue(d *Decoder, val reflect.Value) error {
	n, err := d.DecodeDecimal()
	if err != nil {
		return err
	}

	val.Set(reflect.ValueOf(n))
	return nil
}

func decodeByteArrayValue(d *Decoder, val reflect.Value, size int) error {
	b, err := d.ReadN(size)
	if err != nil {
		return err
	}

	reflect.Copy(val, reflect.ValueOf(b))
	return nil
}
