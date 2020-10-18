package golang

import (
	"github.com/v1990/protoc-gen-goku/descriptors"
	"strconv"
	"strings"
)

type GoPackage struct {
	Name       string
	ImportPath string
}

func (p GoPackage) String() string {
	return p.Name
}

func (p GoPackage) Import() string {
	return p.Name + " " + strconv.Quote(string(p.ImportPath))
}

type GoType struct {
	Base     string
	Pkg      string
	repeated bool
	pointer  bool
}

func (t GoType) String() string {
	return t.defaultFormat()
}

func (t GoType) defaultFormat() string {

	var s []string
	if t.repeated {
		s = append(s, "[]")
	}

	if t.pointer {
		s = append(s, "*")
	}

	if len(t.Pkg) > 0 {
		s = append(s, t.Pkg, ".")
	}

	s = append(s, t.Base)

	return strings.Join(s, "")

}

//
//type GoTypeInterface interface {
//	helper.Formatter
//	fmt.Stringer
//}
//type BaseGoType struct {
//}
//type NestableGoType struct {
//	*BaseGoType
//	d descriptors.Nestable
//}
//type MessageGoType struct {
//	*NestableGoType
//	d *descriptors.DescriptorProto
//}
//type EnumGoType struct {
//	*NestableGoType
//	d *descriptors.EnumDescriptorProto
//}
//type FieldGoType struct {
//	d *descriptors.FieldDescriptorProto
//}
//
//func (c Context) NewFieldGoType(d *descriptors.FieldDescriptorProto) *FieldGoType {
//	return &FieldGoType{d: d}
//}
//
//func (t *FieldGoType) Format(_args ...interface{}) (string, error) {
//	args := helper.Args(_args)
//	switch args.String(0, "default") {
//	case "default":
//		return t.defaultFormat(), nil
//	}
//
//	return "", errors.New("unsupported format")
//
//}
//
//func (t *FieldGoType) String() string {
//	s, err := t.Format()
//	if err != nil {
//		return err.Error()
//	}
//	return s
//}
//
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

//
//func (t *FieldGoType) defaultFormat() string {
//	baseName := t.typeName()
//	if t.d.IsRepeated() {
//		return "[]" + baseName
//	}
//
//	return baseName
//}
//
//func (t *FieldGoType) typeName() string {
//	if s, ok := baseTypes[t.d.GetType()]; ok {
//		return s
//	}
//	return "" // TODO
//}
//
////
////type EnumValueGoType struct {
////	d *descriptors.EnumValueDescriptorProto
////}
////
////type FiledGoType struct {
////	d *descriptors.FieldDescriptorProto
////}
//
//func (c Context) NewNestableGoType(d descriptors.Nestable) *NestableGoType {
//	return &NestableGoType{d: d}
//}
//func (c Context) NewMessageGoType(d *descriptors.DescriptorProto) *MessageGoType {
//	return &MessageGoType{c.NewNestableGoType(d), d}
//}
//
////
//func (c Context) NewEnumGoType(d *descriptors.EnumDescriptorProto) *EnumGoType {
//	return &EnumGoType{c.NewNestableGoType(d), d}
//}
//
////
////func NewEnumValueGoType(d *descriptors.EnumValueDescriptorProto) *EnumValueGoType {
////	return &EnumValueGoType{d: d}
////}
////
////func NewFiledGoType(d *descriptors.FieldDescriptorProto) *FiledGoType {
////	return &FiledGoType{d: d}
////}
//
//func (t *NestableGoType) Format(_args ...interface{}) (string, error) {
//	args := helper.Args(_args)
//	switch args.String(0, "default") {
//	case "default":
//		return t.defaultFormat(), nil
//	}
//
//	return "", errors.New("unsupported format")
//
//}
//
//func (t *NestableGoType) String() string {
//	s, err := t.Format()
//	if err != nil {
//		return err.Error()
//	}
//	return s
//}
//
//func (t *NestableGoType) defaultFormat() string {
//	names := t.TypeNames()
//	return helper.CamelCase(strings.Join(names, "_"))
//}
//
//func (t *NestableGoType) TypeNames() []string {
//	var names []string
//	names = append(names, t.d.GetName())
//	for parent := t.d.ParentMessage(); parent != nil; parent = parent.ParentMessage() {
//		name := parent.GetName()
//		names = append(names, name)
//	}
//	names = helper.ReverseStrings(names)
//	return names
//}
