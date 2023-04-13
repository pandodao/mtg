package mtgpack

import "reflect"

type CustomEncoder interface {
	EncodeMtg(*Encoder) error
}

type CustomDecoder interface {
	DecodeMtg(*Decoder) error
}

func isByteArray(val reflect.Value) (int, bool) {
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() == reflect.Array && val.Type().Elem().Kind() == reflect.Uint8 {
		return val.Len(), true
	}

	return 0, false
}
