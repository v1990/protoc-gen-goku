package golang

import "github.com/v1990/protoc-gen-goku/descriptors"

func IsGoKeyword(s string) bool {
	return goKeywords[s]
}

func IsGoPredeclaredIdentifier(s string) bool {
	return goPredeclaredIdentifier[s]
}

var goKeywords = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

var goPredeclaredIdentifier = map[string]bool{
	"append":     true,
	"bool":       true,
	"byte":       true,
	"cap":        true,
	"close":      true,
	"complex":    true,
	"complex128": true,
	"complex64":  true,
	"copy":       true,
	"delete":     true,
	"error":      true,
	"false":      true,
	"float32":    true,
	"float64":    true,
	"imag":       true,
	"int":        true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"int8":       true,
	"iota":       true,
	"len":        true,
	"make":       true,
	"new":        true,
	"nil":        true,
	"panic":      true,
	"print":      true,
	"println":    true,
	"real":       true,
	"recover":    true,
	"rune":       true,
	"string":     true,
	"true":       true,
	"uint":       true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uint8":      true,
	"uintptr":    true,
}

var baseTypes = map[descriptors.FieldDescriptorProto_Type]string{
	descriptors.FieldDescriptorProto_TYPE_DOUBLE:   "float64",
	descriptors.FieldDescriptorProto_TYPE_FLOAT:    "float",
	descriptors.FieldDescriptorProto_TYPE_INT64:    "int64",
	descriptors.FieldDescriptorProto_TYPE_UINT64:   "uint64",
	descriptors.FieldDescriptorProto_TYPE_INT32:    "int32",
	descriptors.FieldDescriptorProto_TYPE_FIXED64:  "uint64",
	descriptors.FieldDescriptorProto_TYPE_FIXED32:  "int32",
	descriptors.FieldDescriptorProto_TYPE_BOOL:     "bool",
	descriptors.FieldDescriptorProto_TYPE_STRING:   "string",
	descriptors.FieldDescriptorProto_TYPE_BYTES:    "[]byte",
	descriptors.FieldDescriptorProto_TYPE_UINT32:   "uint32",
	descriptors.FieldDescriptorProto_TYPE_SFIXED32: "int32",
	descriptors.FieldDescriptorProto_TYPE_SFIXED64: "int64",
	descriptors.FieldDescriptorProto_TYPE_SINT32:   "int32",
	descriptors.FieldDescriptorProto_TYPE_SINT64:   "int64",
	//descriptors.FieldDescriptorProto_TYPE_ENUM:     "",
	//descriptors.FieldDescriptorProto_TYPE_MESSAGE:  "",
	//descriptors.FieldDescriptorProto_TYPE_GROUP:    "",
}
