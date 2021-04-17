// go/template helper functions
package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/v1990/protoc-gen-goku/helper/dump"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

var Functions = map[string]interface{}{
	"env": GetEnv,

	// base filename with extension. e.g. proto/hello.proto => hello.proto
	"filename": filepath.Base,
	// base filename without extension e.g. proto/hello.proto => hello
	"basename": func(filename string) string {
		return strings.TrimSuffix(
			filepath.Base(filename),
			filepath.Ext(filename),
		)
	},

	"eq":    Eq,
	"in":    In,
	"notIn": NotIn,

	"toString": ToString,
	"showJSON": ShowJSON,

	"contact":     Contact,
	"empty":       Empty,
	"filterEmpty": FilterEmpty,

	"format": Format,

	"trimLinePrefix": TrimLinePrefix,
	"trimLineSuffix": TrimLineSuffix,
	"trimLineLeft":   TrimLineLeft,
	"trimLineRight":  TrimLineRight,
	"withLinePrefix": WithLinePrefix,
	"withLineSuffix": WithLineSuffix,
	"withLineSlash":  WithLineSlash,
	"dump":           dump.Dump,
	"dump2":          dump.Dump2,
}

func RegisterFuncMap(functions map[string]interface{}) {
	for name, f := range functions {
		Functions[name] = f
	}
}

func GetEnv(key string, def ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 && len(def) > 0 {
		return def[0]
	}
	return v
}

// 显示格式化后的json数据
//  depth - 换行的深度
func ShowJSON(data interface{}, depth_ ...int) string {
	// TODO 写的太乱了
	space := " "
	spaceCount := 2
	indent := strings.Repeat(space, spaceCount)
	depth := 0
	if len(depth_) > 0 {
		depth = depth_[0]
		if depth == 0 {
			indent = ""
		}
	}

	var err error
	var content []byte

	if p, ok := data.(proto.Message); ok {
		buf := bytes.NewBuffer(nil)
		marshaler := jsonpb.Marshaler{Indent: indent}
		err = marshaler.Marshal(buf, p)
		content = buf.Bytes()
	} else if len(indent) > 0 {
		content, err = json.MarshalIndent(data, "", indent)
	} else {
		content, err = json.Marshal(data)
	}

	if err != nil {
		return err.Error()
	}

	str := string(content)
	// 没有缩进
	if len(indent) == 0 {
		return str
	}
	numSliceReg := regexp.MustCompile("\n\\s+(\\d+,?)")
	str = numSliceReg.ReplaceAllString(str, "${1}")

	numSliceReg2 := regexp.MustCompile("(\\d+?)\\s*\n\\s*]")
	str = numSliceReg2.ReplaceAllString(str, "${1}]")
	// 没有指定缩进级别
	if len(depth_) == 0 {
		return str
	}

	spaceReg := regexp.MustCompile(fmt.Sprintf("\n {%d,}", spaceCount*depth+1))
	rBraceReg := regexp.MustCompile(fmt.Sprintf("\n {%d,}\\}", spaceCount*depth))
	rBraceReg2 := regexp.MustCompile(fmt.Sprintf("\n {%d,}\\]", spaceCount*depth))

	str = spaceReg.ReplaceAllString(str, "")
	str = rBraceReg.ReplaceAllString(str, "}")
	str = rBraceReg2.ReplaceAllString(str, "]")

	return str
}

func JsonMarshal(v interface{}) []byte {
	s, _ := json.Marshal(v)
	return s
}

func JsonMarshalIndent(v interface{}) []byte {
	s, _ := json.MarshalIndent(v, "  ", "  ")
	return s
}

func In(needle interface{}, haystack ...interface{}) (ok bool) {
	if len(haystack) == 0 {
		return false
	}
	for _, v := range haystack {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice, reflect.Array:
			v2 := reflect.ValueOf(v)
			l := v2.Len()
			for i := 0; i < l; i++ {
				if Eq(needle, v2.Index(i).Interface()) {
					return true
				}
			}
		case reflect.Map:
			v2 := reflect.ValueOf(v)
			for _, keyValue := range v2.MapKeys() {
				if Eq(needle, keyValue.Interface()) {
					return true
				}
			}
		default:
			if Eq(needle, v) {
				return true
			}
		}
	}
	return false
}

func NotIn(needle interface{}, haystack ...interface{}) bool {

	return !In(needle, haystack...)
}

func Eq(a interface{}, b interface{}) (ok bool) {
	av := reflect.ValueOf(a)
	if av.Kind() == reflect.Ptr {
		a = av.Elem().Interface()
	}
	bv := reflect.ValueOf(b)
	if bv.Kind() == reflect.Ptr {
		b = bv.Elem().Interface()
	}

	if reflect.DeepEqual(a, b) {
		return true
	}
	if ToString(a) == ToString(b) {
		return true
	}

	return false
}

func Struct2Map(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return data
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 反转数组
func ReverseStrings(s []string) []string {
	l := len(s)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// CamelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X') // TODO 是否有必要这么处理？
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// CamelCaseSlice is like CamelCase, but the argument is a slice of strings to
// be joined with "_".
func CamelCaseSlice(elem []string) string { return CamelCase(strings.Join(elem, "_")) }

// dottedSlice turns a sliced name into a dotted name.
func dottedSlice(elem []string) string { return strings.Join(elem, ".") }

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func GetName(v interface{}) string {
	switch t := v.(type) {
	case interface{ GetName() string }:
		return t.GetName()
	case interface{ Name() string }:
		return t.Name()
	case string:
		return t
	case *string:
		return *t
	default:
		return fmt.Sprintf("%T", v)
	}
}
