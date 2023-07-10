package mtgpack

import (
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

// CustomEncoder is an interface that defines the Encoder for a custom encoding scheme.
type CustomEncoder interface {
	EncodeMtg(*Encoder) error
}

// CustomDecoder is an interface that defines the Decoder for a custom decoding scheme.
type CustomDecoder interface {
	DecodeMtg(*Decoder) error
}

// isByteArray is a function that checks if the value provided is an array of bytes.
func isByteArray(val reflect.Value) (int, bool) {
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() == reflect.Array && val.Type().Elem().Kind() == reflect.Uint8 {
		return val.Len(), true
	}

	return 0, false
}

var (
	// customDecoderType is a variable that stores a reflection type for CustomDecoder interface
	customDecoderType = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
	customEncoderType = reflect.TypeOf((*CustomEncoder)(nil)).Elem()

	// decimalType is a reflection type for decimal.Decimal
	decimalType = reflect.TypeOf(decimal.Decimal{})

	// timeType is a reflection type for time.Time
	timeType = reflect.TypeOf(time.Time{})
)
