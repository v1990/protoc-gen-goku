package golang

import (
	"github.com/pkg/errors"
	"github.com/v1990/protoc-gen-goku/helper"
	"strconv"
	"strings"
)

type GoPackage struct {
	Name       string // 包 别名
	ImportPath string // 包 全路径
}

func (p GoPackage) String() string {
	return p.Name
}

func (p GoPackage) Import() string {
	return p.Name + " " + strconv.Quote(string(p.ImportPath))
}

type GoType struct {
	name     string     // 基本类型名，
	repeated bool       // 是否数组
	pointer  bool       // 是否指针
	isBase   bool       // 是否基础类型(如int)，如果是，pkg = nil
	pkg      *GoPackage // 包信息
}

func (t *GoType) Package() *GoPackage {
	if t == nil {
		return nil
	}
	return t.pkg
}

func (t *GoType) Name() string {
	if t == nil {
		return ""
	}
	return t.name
}

func (t *GoType) String() string {
	return t.defaultFormat()
}

var _ helper.FormatPipeline = new(GoType)
var _ helper.SetPipeline = new(GoType)

// implemented helper.FormatPipeline
func (t *GoType) FormatPipeline(args helper.Args) (string, error) {
	if args.Len() == 0 || args.String(0, "default") == "default" {
		return t.defaultFormat(), nil
	}
	return "", nil
}

// implemented helper.SetPipeline
func (t *GoType) SetPipeline(field string, value interface{}) interface{} {
	return t.SetField(field, value)
}
func (t *GoType) defaultFormat() string {

	var s []string
	if t.repeated {
		s = append(s, "[]")
	}

	if t.pointer {
		s = append(s, "*")
	}

	if t.pkg != nil {
		s = append(s, t.pkg.Name, ".")
	}

	s = append(s, t.name)

	return strings.Join(s, "")

}

func (t *GoType) Import() string {
	return t.pkg.Import()
}

func (t *GoType) SetField(field string, value interface{}) *GoType {
	switch field {
	case "package":
		return t.SetPackage(value)
	case "pointer":
		return t.SetPointer(value)
	}
	Throws(errors.Errorf("can not set GoType field(%s)=value(%v)", field, value))
	return nil
}

func (t *GoType) SetPackage(value interface{}) *GoType {
	tt := t.Copy()

	if helper.Empty(value) {
		tt.pkg = nil
		return tt
	}

	switch pkg := value.(type) {
	case string:
		tt.pkg = GetGoPackageByName(pkg)
	case *GoPackage:
		tt.pkg = pkg
	}
	if tt.pkg == nil {
		Throws(errors.Errorf("can not set GoType field(%s)=value(%v)", "package", value))
		return tt
	}

	return tt
}

func (t *GoType) Copy() *GoType {
	tt := new(GoType)
	*tt = *t
	return tt
}

func (t *GoType) SetPointer(value interface{}) *GoType {
	tt := t.Copy()
	if helper.Empty(value) {
		tt.pointer = false
		return tt
	}
	b, err := helper.ParseBool(value)
	if err != nil {
		Throws(errors.Errorf("SetPointer: %w", err))
		return tt
	}
	tt.pointer = b

	return tt

}

func Throws(err error) {
	helper.Throws(err)
}
