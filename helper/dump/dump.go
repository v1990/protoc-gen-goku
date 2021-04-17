package dump

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

//
//import (
//	"bytes"
//	"fmt"
//	"reflect"
//	"runtime/debug"
//	"strings"
//)
//

func Dump2(v interface{}) string {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("panic: %s \n", r)
	//		debug.PrintStack()
	//	}
	//}()

	m := &marshaler{
		indent:   0,
		maxDepth: 5,
	}
	m.dump(v)
	return m.buf.String()
}

//
type marshaler struct {
	buf      bytes.Buffer
	indent   int
	maxDepth int
}

var _id = 0

func (m *marshaler) dump(in interface{}) {

	m.indent++
	defer func() {
		m.indent--
	}()

	var val reflect.Value
	switch in := in.(type) {
	case reflect.Value:
		val = in
	default:
		val = reflect.ValueOf(in)
	}
	elem := reflect.Indirect(val)

	if m.indent >= m.maxDepth {
		m.writeln(fmt.Sprintf("%#v", val.Interface()))
		return
	}

	_ = in // don't use in

	if val.IsValid() && val.CanInterface() && val.Kind() != reflect.Invalid {
		m.writeType(val)

		typ := val.Type()
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				m.writeln("(nil)")
				return
			}
			typ = elem.Type()
		}
		_id++
		id := fmt.Sprintf(" // [%d] %s %T", _id, typ.Kind(), val.Interface())

		switch typ.Kind() {
		//case reflect.Ptr: // no this cace
		case reflect.Array, reflect.Slice:
			l := val.Len()
			m.writeln("{", id)
			for i := 0; i < l; i++ {
				m.dump(val.Index(i))
			}
			m.writeln("}", id)
		case reflect.Struct:
			m.writeln("{", id)
			n := typ.NumField()
			for i := 0; i < n; i++ {
				ft := typ.Field(i)
				m.writeIndent()
				m.write(ft.Name, ":", " ")
				m.dump(elem.Field(i))
			}
			m.writeIndent(m.indent - 1)
			m.writeln("}", id)
		case reflect.Map: // TODO
			m.writeln("{", id)

			keys := val.MapKeys()
			for _, key := range keys {
				m.writeIndent()
				m.write(fmt.Sprintf("%v: ", key.Interface()))
				m.dump(val.MapIndex(key))
			}
			m.writeln("}", id)
		case reflect.Interface: // TODO
			m.dump(val)
			//m.writeln(fmt.Sprintf("(%#v)", val.Interface()))
		case reflect.Func:
			m.writeln(fmt.Sprintf("(%#v)", val.Interface()))
		default:
			m.writeln(fmt.Sprintf("(%#v)", val.Interface()))
		}
	} else {
		m.write("<invalid>")
	}
}
func (m *marshaler) write(s ...string) {
	m.buf.WriteString(strings.Join(s, ""))
}
func (m *marshaler) writeln(s ...string) {
	m.buf.WriteString(strings.Join(s, ""))
	m.buf.WriteString("\n")
}
func (m *marshaler) writeIndent(indent ...int) {
	l := m.indent
	if len(indent) > 0 {
		l = indent[0]
	}
	for i := 0; i < l; i++ {
		m.buf.WriteString("-")
	}
}
func (m *marshaler) writeType(val reflect.Value) {
	m.buf.WriteString(fmt.Sprintf("%T", val.Interface()))
}

//
type Exporter interface {
	ExportFields() map[string]interface{}
}

//
////
////func dumpValue(v interface{})interface{}{
////
////}
////
