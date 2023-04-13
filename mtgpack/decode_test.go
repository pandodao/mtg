package mtgpack

import (
	"reflect"
	"testing"
)

func TestDecodeUUID(t *testing.T) {
	var b = make([]byte, 32)

	// typ := reflect.TypeOf(b)
	val := reflect.ValueOf(&b)

	if val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	t.Logf("is array: %t", val.Kind() == reflect.Array)

	if val.Kind() == reflect.Array {
		t.Log("elem type is", val.Type().Elem())
	}

	t.Logf("len: %d", val.Len())
}
