package golang

import (
	"github.com/pkg/errors"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"strconv"
	"strings"
)

var (
	globalPackages = map[string]*GoPackage{}
)

// golang 对context提供的处理函数集合
// 其他基于 golang 的插件可以继承此类进行扩展
type Context struct {
	*goku.Context
	markedImports map[string]*GoPackage
}

func NewContext(ctx *goku.Context) *Context {
	return &Context{
		Context:       ctx,
		markedImports: make(map[string]*GoPackage),
	}
}

func (c Context) FuncMap() goku.FuncMap {
	return goku.FuncMap{
		"GoType":             c.GoType,
		"GoPackage":          c.GoPackage,
		"GoTypeBase":         c.GoTypeBase,
		"GoComments":         c.GoComments,
		"GoImport":           c.GoImport,
		"GoImportDependency": c.GoImportDependency,
		"GoMarkImport":       c.GoMarkImport,
	}
}

// interface To GoType
func (c Context) GoType(value interface{}) *GoType {
	switch v := value.(type) {
	case *GoType:
		return v
	case descriptors.ProtoType:
		return c.PbType2GoType(v)
	case *descriptors.FieldDescriptorProto:
		return c.PbField2GoType(v)
	case string: // e.g. method.inputType/outputType
		if p := c.GetDescriptorByName(v); p != nil {
			return c.PbType2GoType(p)
		}
	}
	c.ThrowsOnErr(errors.Errorf("GoType: unknown type:%T name:%s", value, helper.GetName(value)))
	return nil
}

// 返回基本类型名
// []*a.T => T
func (c Context) GoTypeBase(v interface{}) string {
	return c.GoType(v).Name()
}

func (c Context) GoPackage(value interface{}) *GoPackage {
	switch v := value.(type) {
	case *GoPackage:
		return v
	case *GoType:
		return v.pkg
	case string:
		return GetGoPackageByName(v)
	case *descriptors.FileDescriptorProto:
		return c.FileGoPackage(v)
	case descriptors.DescriptorCommon:
		if f := v.File(); !f.Empty() {
			return c.FileGoPackage(v.File())
		}
	}
	return nil
}
func GetGoPackageByName(name string) *GoPackage {
	if p, ok := globalPackages[name]; ok {
		return p
	}
	return RegisterPackage(name, "")
}

func (c Context) RegisterPackage(pkg string, name string) *GoPackage {
	return RegisterPackage(pkg, name)
}
func RegisterPackage(pkg string, name string) *GoPackage {
	// 已注册，直接返回
	if p := GetGoPackageByPath(pkg); p != nil {
		return p
	}
	// 提取 包别名
	if len(name) == 0 {
		// foo/bar => bar
		name = helper.BaseName(pkg)
	}

	if IsGoKeyword(name) || IsGoPredeclaredIdentifier(name) {
		name = name + "_"
	}

	// 确保包别名唯一
	for i, orig := 0, name; globalPackages[name] != nil; i++ {
		name = orig + "_" + strconv.Itoa(i)
	}

	p := &GoPackage{
		Name:       name,
		ImportPath: pkg,
	}

	globalPackages[p.Name] = p
	return p

}

func GetGoPackageByPath(importPath string) *GoPackage {
	for _, p := range globalPackages {
		if p.ImportPath == importPath {
			return p
		}
	}
	return nil
}

// File to GoPackage
func (c Context) FileGoPackage(file *descriptors.FileDescriptorProto) *GoPackage {
	var name, pkg string

	if pkg = file.GetOptions().GetGoPackage(); len(pkg) > 0 {
		// 如：option go_package = "github.com/golang/protobuf/protoc-Generator-go/descriptor;descriptor";
		if n := strings.LastIndex(pkg, ";"); n > 0 {
			name = pkg[n+1:]
			pkg = pkg[:n]
		}
	} else if pkg = file.GetPackage(); len(pkg) > 0 {
		// TODO 使用 proto package 作为 go package 是不规范的，
		//		考虑给外部提供 包名映射的方法（提供在配置的 Data）
		// 如： package google.protobuf;
		pkg = strings.Trim(pkg, ".")
		pkg = strings.ReplaceAll(pkg, ".", "/")
	} else {
		c.Fatal("can not parse file go package info. for: %s", file.GetName())
	}

	if name == "" {
		if n := strings.LastIndex(pkg, "/"); n > 0 {
			name = pkg[n+1:]
		} else {
			name = pkg
		}
	}

	return c.RegisterPackage(pkg, name)
}

// Field to GoType
func (c Context) PbField2GoType(field *descriptors.FieldDescriptorProto) *GoType {

	t := &GoType{}
	if s, ok := baseTypes[field.GetType()]; ok {
		t.name = s
		t.isBase = true
	} else {
		// Message / Enum...
		// TODO OneOf / Group 等类型处理
		pbType := c.MustGetDescriptorByName(field.GetTypeName())
		t2 := c.PbType2GoType(pbType)
		*t = *t2 // copy
		switch field.GetType() {
		case descriptors.FieldDescriptorProto_TYPE_MESSAGE:
			t.pointer = true
		}
	}

	// 是否数组
	t.repeated = field.IsRepeated()

	return t
}

// ProtoType to GoType
func (c Context) PbType2GoType(pbType descriptors.ProtoType) *GoType {
	info := pbType.ProtoType()
	base := helper.CamelCaseSlice(info.TypeNames())
	pkg := c.FileGoPackage(pbType.File())

	pointer := false
	switch pbType.(type) {
	case *descriptors.DescriptorProto: // message 默认为指针
		pointer = true
	case *descriptors.EnumDescriptorProto: // enum 默认不需要指针
		pointer = false
	}

	return &GoType{
		name:     base,
		repeated: false,
		pointer:  pointer,
		pkg:      pkg,
	}
}

// GoComments returns golang style comments string
//  supported arg:
//  - string
//  - []string
//  - descriptors.DescriptorCommon
//  - descriptors.SourceCodeInfo_Location
func (c Context) GoComments(arg interface{}) string {
	switch v := arg.(type) {
	case string:
		return toGoComments(v)
	case []string:
		return c.GoComments(strings.Join(v, "\n"))
	case *descriptors.SourceCodeInfo_Location:
		var comments []string
		// 顶部独立注释
		comments = append(comments, v.GetLeadingDetachedComments()...)
		// 顶部注释
		comments = append(comments, v.GetLeadingComments())
		// 行尾注释
		comments = append(comments, v.GetTrailingComments())
		return c.GoComments(comments)
	case descriptors.DescriptorCommon:
		return c.GoComments(v.Comments())
	}
	return c.GoComments(helper.ToString(arg))

}

func toGoComments(src string) string {
	return helper.WithLineSlash(src)
}

// GoImport returns golang style import line string.
func (c Context) GoImport(arg interface{}) string {
	if arg == nil {
		return ""
	}
	switch v := arg.(type) {
	case *GoPackage:
		return v.Import()
	case *GoType:
		return c.GoImport(v.pkg) // goto case *GoPackage
	case *descriptors.FieldDescriptorProto:
		return c.GoImport(c.PbField2GoType(v)) // goto case *GoType
	case *descriptors.FileDescriptorProto:
		return c.GoImport(c.FileGoPackage(v)) // goto case *GoPackage
	case descriptors.ProtoType:
		return c.GoImport(c.PbType2GoType(v)) // goto case *GoType
	}
	c.ThrowsOnErr(errors.Errorf("failed to parse import: %#v", arg))
	return ""
}

// GoMarkImport 标记导入 【v】 所在的包，并在最终生成代码时加入到 import 中
//    v: 能转化为 GoPackage 的对象，并且没有对【v】做任何修改，直接返回.
// 所以 GoMarkImport 可以放在 pipeline 的任意合适的位置(前一个pipeline必须返回【v】)
//    而 GoImport 必须放在 import 括号之中（直接生成了字符串）
//    还有 GoImportDependency 也必须放在 import 括号之中,但是副作用很大，
// 因为会导入proto文件中声明的所有 import，虽然最后生成代码时会自动 clean unused import.
//
//  Example:
//		{{$Input:=$Method.InputType|GoType|GoMarkImport}}
func (c Context) GoMarkImport(v interface{}) interface{} {
	p := c.GoPackage(v)
	c.markedImports[p.Name] = p
	return v
}

// GoImportDependency returns golang style import lines for file Dependency
func (c Context) GoImportDependency(arg interface{}) string {
	switch v := arg.(type) {
	case *descriptors.FileDescriptorProto:
		var lines []string
		for _, d := range v.GetDependency() {
			f := c.GetFileByName(d)
			//log.Println("GetDependency:",d,f==nil)
			line := c.GoImport(f)
			lines = append(lines, line)
		}
		return strings.Join(lines, "\n")
	case descriptors.DescriptorCommon:
		return c.GoImportDependency(v.File()) // goto case *descriptors.FileDescriptorProto
	}

	c.ThrowsOnErr(errors.Errorf("GoImportDependency: unsupport arg: %T", arg))
	return ""
}
