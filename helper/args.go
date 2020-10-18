package helper

import "reflect"

type Args []interface{}

func (args Args) Get(index int) (v interface{}, ok bool) {
	if index < 0 {
		return nil, false
	}
	if len(args) <= index {
		return nil, false
	}
	return args[index], true
}
func (args Args) Def(index int, def interface{}) interface{} {
	if v, ok := args.Get(index); ok {
		return v
	}
	return def
}

// dereference 是否解引用，如果是 reflect.Ptr 则返回其原始类型，如 *int(ptr) => int
func (args Args) Kind(index int, dereference bool) reflect.Kind {
	if v, ok := args.Get(index); ok {
		return Kind(v, dereference)
	}
	return reflect.Invalid
}
func (args Args) String(index int, def string) string {
	return ToString(args.Def(index, def))
}
func (args Args) Int(index int, def int) int {
	return ToInt(args.Def(index, def), def)
}
func (args Args) Bool(index int, def bool) bool {
	return ToBool(args.Def(index, def), def)
}

func (args Args) Len() int { return len(args) }

func (args Args) First() interface{} {
	return args.Def(0, nil)
}
func (args Args) Last() interface{} {
	return args.Def(args.Len()-1, nil)
}

func (args *Args) Pop() interface{} {
	if v, ok := args.Get(args.Len() - 1); ok {
		*args = (*args)[:args.Len()-1]
		return v
	}
	return nil
}

func (args *Args) Push(other ...interface{}) Args {
	*args = append(*args, other...)
	return *args
}

func (args *Args) Shift() interface{} {
	if v, ok := args.Get(0); ok {
		*args = (*args)[1:]
		return v
	}
	return nil
}

func (args Args) Flatten() Args {
	var res Args
	for i := 0; i < args.Len(); i++ {
		v := reflect.ValueOf(args[i])
		switch kind(v, true) {
		case reflect.Slice, reflect.Array:
			l := v.Len()
			for j := 0; j < l; j++ {
				res = append(res, v.Index(j).Interface())
			}
		default:
			res = append(res, args[i])
		}
	}

	return res
}

func Kind(v interface{}, dereference_ ...bool) reflect.Kind {
	return kind(reflect.ValueOf(v), dereference_...)
}

func kind(vv reflect.Value, dereference_ ...bool) reflect.Kind {
	var dereference = false
	if len(dereference_) > 0 {
		dereference = dereference_[0]
	}
	kind := vv.Kind()
	if dereference && kind == reflect.Ptr {
		return vv.Elem().Kind()
	}
	return kind
}
