package golang

import (
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"go/format"
	"strconv"
)

type goPlugin struct {
	gen *goku.Generator
}

func init() {
	goku.RegisterPlugin(new(goPlugin))
}

func (p *goPlugin) Name() string {
	return "golang"
}

func (p *goPlugin) Init(g *goku.Generator) {
	p.gen = g

	//p.baseFuncMap = goku.FuncMap{}

}

func (p *goPlugin) BeforeExecute(ctx *goku.Context) {
	c := NewContext(ctx)
	ctx.MergeFuncMap(c.FuncMap())

	//log.Println("goPlugin.Context:", ctx.ShowFunc())

	//ctx.MergeFuncMap(goku.FuncMap{
	//	////"getGoTypeName": p.typeName,
	//	//"NewGoField": func(field *descriptors.FieldDescriptorProto) *GoField {
	//	//	return p.NewGoField(ctx, field.PbDescriptor())
	//	//},
	//	//"GoTypeInterface": func(name string) GoTypeInterface {
	//	//	return p.ToGoType(ctx, name)
	//	//},
	//})

}

func (p *goPlugin) NewFuncMap(ctx *goku.Context) goku.FuncMap {
	f := &Context{
		ctx: ctx,
	}

	return f.FuncMap()
}

//
//type GoField struct {
//	ctx  *goku.Context
//	desc *descriptor.FieldDescriptorProto
//	Type string
//	Zero string
//	Wire string
//	//Start     string
//	//StartType string
//}
//
//func (p *goPlugin) NewGoField(ctx *goku.Context, field *descriptor.FieldDescriptorProto) *GoField {
//	f := &GoField{ctx: ctx, desc: field}
//	typ, wire, zero := p.filedType(ctx, field)
//	f.Type = typ
//	f.Wire = wire
//	f.Zero = zero
//	return f
//}
//
//func (p *goPlugin) filedType(ctx *goku.Context, field *descriptor.FieldDescriptorProto) (typ string, wire string, zero string) {
//	//defer ctx.Recover(nil)
//
//	switch *field.Type {
//	//case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
//	//	typ = "*" + p.ToGoType(ctx, field.GetTypeName()).String()
//	//	zero = "nil"
//	//	wire = "bytes"
//	//case descriptor.FieldDescriptorProto_TYPE_ENUM:
//	//	typ = p.ToGoType(ctx, field.GetTypeName()).String()
//	//	wire = "varint"
//	//	zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_GROUP:
//		// TODO
//		//desc := g.ObjectNamed(field.GetTypeName())
//		//typ, wire = "*"+g.GoTypeName(desc), "group"
//	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
//		typ = "float64"
//		wire = "fixed64"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
//		typ = "float32"
//		wire = "fixed32"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_INT64:
//		typ = "int64"
//		wire = "varint"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_UINT64:
//		typ = "uint64"
//		wire = "varint"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_INT32:
//		typ = "int32"
//		wire = "varint"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_UINT32:
//		typ = "uint32"
//		wire = "varint"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
//		typ = "uint64"
//		wire = "fixed64"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
//		typ = "uint32"
//		wire = "fixed32"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_BOOL:
//		typ = "bool"
//		wire = "varint"
//		zero = "false"
//	case descriptor.FieldDescriptorProto_TYPE_STRING:
//		typ = "string"
//		wire = "bytes"
//		zero = `""`
//	case descriptor.FieldDescriptorProto_TYPE_BYTES:
//		typ = "[]byte"
//		wire = "bytes"
//		zero = "nil"
//	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
//		typ = "int32"
//		wire = "fixed32"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
//		typ = "int64"
//		wire = "fixed64"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_SINT32:
//		typ = "int32"
//		wire = "zigzag32"
//		zero = "0"
//	case descriptor.FieldDescriptorProto_TYPE_SINT64:
//		typ = "int64"
//		wire = "zigzag64"
//		zero = "0"
//	default:
//		// TODO
//		ctx.Warn("unknown type for: %s", field.GetName())
//	}
//
//	if field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
//		typ = "[]" + typ
//		zero = "nil"
//	}
//
//	return
//}

//func (f GoField) IsPtr() bool {
//	if f.ctx.File().GetSyntax() == "proto3" {
//		return false
//	}
//
//	switch f.desc.GetType() {
//	case descriptor.FieldDescriptorProto_TYPE_GROUP:
//		return false
//	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
//		return false
//	case descriptor.FieldDescriptorProto_TYPE_BYTES:
//		return false
//	}
//	return true
//}

func (p *goPlugin) BeforeOut(ctx *goku.Context) {
	content, err := p.format(ctx.Content())
	if err != nil {
		p.gen.Warn("[golang]format err:%s \n %s \n", err.Error(),
			string(p.lineIndex(ctx.Content())))
		return
	}

	ctx.SetContent(content)
}

func (p *goPlugin) format(content []byte) ([]byte, error) {
	source, err := format.Source(content)
	return source, err
}

// 为每行代码开头加上行号
func (p *goPlugin) lineIndex(content []byte) []byte {
	i := 0
	body := helper.EachLine(string(content), func(line string) string {
		i++
		return strconv.Itoa(i) + "\t" + line
	})
	return []byte(body)
}

//
//func (p *goPlugin) ToGoType(ctx *goku.Context, name string) (typeInfo GoTypeInterface) {
//	p.gen.RecordTypeUse(name)
//	typ := p.gen.TypeName(p.gen.ObjectNamed(name))
//	n := strings.Index(typ, ".")
//
//	if n < 0 { // 没有包名，则为包内声明的类型
//		// 如果配置声明：生成文件与proto的生成文件不在同一个包
//		if !helper.ToBool(ctx.Value(FileGoPackageSame)) {
//			typeInfo.Base = typ
//			typeInfo.Pkg.Base = helper.ToString(ctx.Value(FileGoPackageName))
//			typeInfo.Pkg.Path = helper.ToString(ctx.Value(FileGoPackagePath))
//		} else {
//			typeInfo.Base = typ
//		}
//	} else {
//		typeInfo.Pkg.Base = typ[:n]
//		typeInfo.Base = typ[n+1:]
//	}
//
//	return typeInfo
//}
//func (p *goPlugin) parsePackage(file *goku.File) (name, path string) {
//	//pkg := file.GetOptions().GetGoPackage()
//	pkg := file.Options().GoPackage()
//	arr := strings.Split(pkg, ";")
//
//	path = arr[0]
//	name = filepath.Base(path)
//
//	if len(arr) == 2 {
//		name = arr[1]
//	}
//
//	return name, path
//
//}
