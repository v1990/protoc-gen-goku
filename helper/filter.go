package helper

import "reflect"

func FilterEmpty(v ...interface{}) []interface{} {
	var res []interface{}

	args := Contact(v...)
	for _, arg := range args {
		if !Empty(arg) {
			res = append(res, arg)
		}
	}

	return res
}

func Empty(given interface{}) bool {
	if given == nil {
		return true
	}

	if v, ok := given.(interface{ Empty() bool }); ok {
		return v.Empty()
	}

	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return true
	}

	// Basically adapted from text/template.isTrue
	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return !g.Bool()
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
}

func Contact(args_ ...interface{}) []interface{} {
	return Args(args_).Flatten()
}
