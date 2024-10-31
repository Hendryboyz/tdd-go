package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := range val.NumField() {
			walkValue(val.Field((i)))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFunResult := val.Call(nil)
		for _, res := range valFunResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		// use Elem() to extract the underlying value in pointer
		val = val.Elem()
	}
	return val
}
