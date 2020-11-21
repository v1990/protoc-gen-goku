package helper

import (
	"fmt"
	"strconv"
)

func ToString(v interface{}, def ...string) string {
	if v == nil {
		if len(def) > 0 {
			return def[0]
		}
		return ""
	}

	switch v := v.(type) {
	case *string:
		return *v
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		if len(def) > 0 {
			return def[0]
		}
		return fmt.Sprintf("%v", v)
	}
}

func ToBool(v interface{}, _def ...bool) bool {
	def := false
	if len(_def) > 0 {
		def = _def[0]
	}

	switch t := v.(type) {
	case bool:
		return t
	case *bool:
		return *t
	default:
		str := ToString(def, "false")
		if tt, err := strconv.ParseBool(str); err == nil {
			return tt
		}
	}

	return def
}

func ToInt(v interface{}, def int) int {
	// TODO 优化
	vv, err := strconv.Atoi(ToString(v))
	if err != nil {
		return def
	}
	return vv
}

func ParseBool(v interface{}) (bool, error) {
	switch b := v.(type) {
	case bool:
		return b, nil
	case *bool:
		return *b, nil
	case string:
		return strconv.ParseBool(b)
	}
	return false, fmt.Errorf("parseBool: %v", v)
}
