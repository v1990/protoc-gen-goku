package helper

import (
	"fmt"
	"reflect"
)

func init() {
	regFuncStruct(new(helper))
}

type helper struct {
}

func (helper) Help() string {
	return "Help"
}

func regFuncStruct(s interface{}) {
	funcMap := make(map[string]interface{})
	ft := reflect.TypeOf(s)
	fv := reflect.ValueOf(s)
	n := ft.NumMethod()
	for i := 0; i < n; i++ {
		mt := ft.Method(i)
		funcName := mt.Name
		//f := mt.Func.Interface()
		f := fv.MethodByName(funcName).Interface()
		funcMap[funcName] = f
		funcMap[funcName+"Usage"] = func() string {
			return fmt.Sprintf("%#v", mt)
		}

	}

	RegisterFuncMap(funcMap)
}
