package dump

import (
	"encoding/json"
	"reflect"
)

func Dump(v interface{}) string {
	ret, err := json.MarshalIndent(ToMap(v), " ", " ")
	if err != nil {
		panic(err)
	}
	return string(ret)
}

type Map map[string]interface{}

func NewMap() Map {

	return make(map[string]interface{})
}

func (m Map) Set(name string, v interface{}) {
	m[name] = v
}
func (m Map) Copy() Map {
	m2 := NewMap()
	m2.DoMerge(m)
	return m2
}
func (m Map) Merge(mm ...Map) Map {
	m2 := m.Copy()
	m2.DoMerge(mm...)
	return m2
}
func (m Map) DoMerge(mm ...Map) {
	for _, m2 := range mm {
		for k, v := range m2 {
			m[k] = v
		}
	}
}

type Int int

func (i *Int) Incr(n int) { *i += Int(n) }
func (i *Int) Decr(n int) { *i -= Int(n) }

func (m Map) Contains(name string) bool {
	_, ok := m[name]
	return ok
}
func ToMap(v interface{}) Map {
	m := NewMap()
	if e, ok := v.(Exporter); ok {
		for k, v := range e.ExportFields() {
			m.Set(k, dumpAny(v))
		}
	}

	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return m
	}
	elem := reflect.Indirect(val)
	if !val.IsValid() {
		return m
	}
	typ := elem.Type()
	if typ.Kind() != reflect.Struct {
		return m
	}

	for i := 0; i < typ.NumField(); i++ {
		fv := elem.Field(i)
		ft := typ.Field(i)
		if !fv.CanInterface() || !fv.IsValid() {
			continue
		}

		// ExportFields 已经导出的，跳过
		if m.Contains(ft.Name) {
			continue
		}

		// 匿名字段,可以被直接访问到的
		if ft.Anonymous {
			m.DoMerge(ToMap(fv.Interface()))
		}
		m.Set(ft.Name, dumpAny(fv.Interface()))

	}

	return m
}

func dumpAny(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if !val.IsValid() || !val.CanInterface() {
		return v
	}
	elem := reflect.Indirect(val)
	if !elem.IsValid() {
		return v
	}
	typ := elem.Type()

	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		return dumpSlice(elem)
	case reflect.Struct:
		return ToMap(v)
	case reflect.Map:
		return v // TODO
	default:
		return v
	}
}

func dumpSlice(v reflect.Value) []interface{} {
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = dumpAny(v.Index(i).Interface())
	}
	return ret
}
