/*

此包是为扩展 google.protobuf.descriptor 的功能，
所有结构除了 protoc-gen-go 为每个 Field 生成的 getter: GetField() 之外，
还生成了 Field() 方法。

因为在 go/template 中，直接使用 {{.Message.Field}} 可能会有空指针问题，
而使用 {{.Message.GetField}} 又觉得太丑。
所以，现在可以安全又愉快地使用 {{.Message.Field}}


此外，还扩展了一些功能，参考：
	type DescriptorCommon interface{} 全体都有的扩展方法
	type Nestable interface{}  可嵌套的对象： message/enum
	type ProtoType interface{} proto类型信息：  message/enum
*/
package descriptors

import (
	"fmt"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/v1990/protoc-gen-goku/helper"
	"log"
	"strings"
)

//////////////////////////////////////////////////////////////////////
// NewFileDescriptorProto 是唯一允许公开的构造器
func NewFileDescriptorProto(file *descriptor.FileDescriptorProto) *FileDescriptorProto {
	t := newFileDescriptorProto(file)
	//t.setFile(t)
	return t
}

// 每个结构都有的方法
type DescriptorCommon interface {
	// 当前对象是否为空
	Empty() bool

	// 当前对象所属文件
	File() *FileDescriptorProto

	// 上级对象
	//  除了 FileDescriptorProto ，其他对象的 Parent 都不为 nil
	//  如 Method 的上级必定为 Service
	//
	//  注意：top level 的 message/enum 的 Parent 为 File
	//	    而嵌套(nested)的 message/enum 的 Parent 则为 message
	Parent() DescriptorCommon

	// 索引值
	//  如 Field 在其 Message 的序号
	//  默认为 -1
	Index() int

	// 获取source中的注释信息
	Comments() *SourceCodeInfo_Location
	// 获取 SourceCodeInfo.Location .Path
	LocationPath() LocationPath
	// 是否为顶级的消息
	//	即 Parent = FileDescriptorProto
	IsTopLevel() bool
}

type NamedDescriptor interface {
	GetName() string
	DescriptorCommon
}

// 可嵌套在message中的结构
//  See DescriptorProto / EnumDescriptorProto
//  See NestableDescriptors
type Nestable interface {
	NamedDescriptor

	// 是否嵌套的
	IsNested() bool
	// 所属的message
	ParentMessage() *DescriptorProto
}

// proto的基本类型信息. 目前有以下struct实现了此interface
// - message: DescriptorProto
// - enum: EnumDescriptorProto
// See ProtoTypeDescriptors
type ProtoType interface {
	NamedDescriptor
	// proto type 信息
	ProtoType() *PbTypeInfo
}

type common struct {
	// 附属的主体,肯定不会为nil
	desc   DescriptorCommon
	file   *FileDescriptorProto
	index  *int
	parent DescriptorCommon
}

func (t *common) setDescriptor(desc DescriptorCommon) {
	t.desc = desc
}

////////////////////////////////////////////////////////////////////////

// implement DescriptorCommon.Index()
//func (t *common) Index() int { return t.getIndex() }
func (t *common) getIndex() int {
	if t == nil {
		return -1
	}
	if t.index == nil {
		return -1
	}
	return *t.index
}
func (t *common) setIndex(i int) {
	t.index = &i
}

////////////////////////////////////////////////////////////////////////

// implement DescriptorCommon.File()
//func (t *common) File() *FileDescriptorProto    { return t.getFile() }
func (t *common) getFile() *FileDescriptorProto { return t.file }
func (t *common) setFile(file *FileDescriptorProto) {
	t.file = file
}

////////////////////////////////////////////////////////////////////////

// implement DescriptorCommon.Parent()
//func (t *common) Parent() DescriptorCommon    { return t.getParent() }
func (t *common) getParent() DescriptorCommon { return t.parent }
func (t *common) setParent(parent DescriptorCommon) {
	if parent == nil {
		return
	}
	t.parent = parent
	if f, ok := parent.(*FileDescriptorProto); ok {
		t.setFile(f)
	} else {
		t.setFile(parent.File())
	}
}

////////////////////////////////////////////////////////////////////////

// implement DescriptorCommon.IsTopLevel
func (t *common) IsTopLevel() bool { return t.isTopLevel() }
func (t *common) isTopLevel() bool {
	if _, ok := t.desc.Parent().(*FileDescriptorProto); ok {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////

//// implement Nestable.IsNested()
//func (t *DescriptorProto) IsNested() bool { return t.isNested() }
//
//// implement Nestable.IsNested()
//func (t *EnumDescriptorProto) IsNested() bool { return t.isNested() }

// 是否嵌套的在 message 中
//  see DescriptorProto.NestedType
//  see DescriptorProto.EnumType
func (t *common) isNested() bool { return t.parentMessage() != nil }

////////////////////////////////////////////////////////////////////////

func (t *common) parentMessage() *DescriptorProto {
	if v, ok := t.desc.Parent().(*DescriptorProto); ok && v != nil {
		return v
	}
	return nil
}

////////////////////////////////////////////////////////////////////////

type LocationPath []int32

// implement DescriptorCommon.LocationPath()
//func (t *common) LocationPath() LocationPath {
//	return t.getLocationPath()
//}

func (t *common) getLocationPath() LocationPath {
	// path规则 : [ parentPath... , tag  [,index] ]
	var path LocationPath

	if !t.IsTopLevel() {
		path = append(path, t.desc.Parent().LocationPath()...)
	}

	if tag := t.getPathTag(); tag >= 0 {
		path = append(path, tag)
	}
	index := t.desc.Index()
	if index >= 0 {
		path = append(path, int32(index))
	}
	return path
}

func (t *common) getPathTag() int32 {
	const (
		// tag numbers in FileDescriptorProto (top level)
		packagePath = 2 // package
		messagePath = 4 // message_type
		enumPath    = 5 // enum_type
		servicePath = 6 // service
		// tag numbers in DescriptorProto
		messageFieldPath   = 2 // field
		messageMessagePath = 3 // nested_type
		messageEnumPath    = 4 // enum_type
		messageOneofPath   = 8 // oneof_decl
		// tag numbers in EnumDescriptorProto
		enumValuePath = 2 // value
		// tag numbers in ServiceDescriptorProto
		serviceMethodPath = 2
	)

	switch obj := t.desc.(type) {
	case *DescriptorProto: // message
		if obj.IsNested() {
			return messageMessagePath
		}
		return messagePath
	case *EnumDescriptorProto:
		if obj.IsNested() {
			return messageEnumPath
		}
		return enumPath
	case *FieldDescriptorProto:
		return messageFieldPath
	case *ServiceDescriptorProto:
		return servicePath
	case *MethodDescriptorProto:
		return serviceMethodPath
	case *EnumValueDescriptorProto:
		return enumValuePath
	case *OneofDescriptorProto:
		return messageOneofPath
	case *FileDescriptorProto:
		// return packagePath ??
		return -1
	default:
		log.Printf("[warn]path tag not found: %T\n", t.desc)
		return -1
	}
}

func (p LocationPath) String() string {
	return fmt.Sprintf("%v", []int32(p))
}

////////////////////////////////////////////////////////////////////////

// implement DescriptorCommon.Comments()
func (t *common) getComments() *SourceCodeInfo_Location {
	return t.desc.File().findLocationByPath(t.getLocationPath())
}

// 默认的注释
func (t *SourceCodeInfo_Location) String() string {
	if t.Empty() {
		return ""
	}
	return t.GetLeadingComments()
}

//////////////////////////////////////////////////////////////////////
// FileDescriptorProto Extensions

func (t *FileDescriptorProto) findLocationByPath(path LocationPath) *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	if t.locByPath == nil {
		t.locByPath = make(map[string]*SourceCodeInfo_Location)
		for _, loc := range t.GetSourceCodeInfo().GetLocation() {
			pathStr := LocationPath(loc.GetPath()).String()
			t.locByPath[pathStr] = loc
		}
		//log.Println("locations: ", helper.ShowJSON(t.locByPath, 1))
	}

	return t.locByPath[path.String()]
}
func (t *FileDescriptorProto) IsProto2() bool {
	return t.GetSyntax() == "proto2"
}
func (t *FileDescriptorProto) IsProto3() bool {
	return t.GetSyntax() == "proto3"
}

//////////////////////////////////////////////////////////////////////
// MethodDescriptorProto Extensions

// Method 所属的 Service
func (t *MethodDescriptorProto) Service() *ServiceDescriptorProto {
	if t.Empty() {
		return nil
	}
	if s, ok := t.parent.(*ServiceDescriptorProto); ok && s != nil {
		return s
	}
	return nil
}

//////////////////////////////////////////////////////////////////////
// FieldDescriptorProto Extensions

func (t *FieldDescriptorProto) IsRequired() bool {
	return t.GetLabel() == FieldDescriptorProto_LABEL_REQUIRED
}

func (t *FieldDescriptorProto) IsOptional() bool {
	return t.GetLabel() == FieldDescriptorProto_LABEL_OPTIONAL
}

func (t *FieldDescriptorProto) IsRepeated() bool {
	return t.GetLabel() == FieldDescriptorProto_LABEL_REPEATED
}

//////////////////////////////////////////////////////////////////////

type PbTypeInfo struct {
	d     NamedDescriptor
	names []string
}

func (t *PbTypeInfo) TypeNames() []string {
	return t.names
}

// FullTypeName 完整的类型名称,如：.g
func (t *PbTypeInfo) FullTypeName() string {
	dottedPkg := "." + t.d.File().GetPackage()
	if dottedPkg != "." {
		dottedPkg += "."
	}
	return dottedPkg + strings.Join(t.names, ".")
}

func (t PbTypeInfo) Package() string {
	return t.d.File().GetPackage()
}

func nestedTypeNames(d Nestable) []string {
	var names []string
	names = append(names, d.GetName())
	for parent := d.ParentMessage(); parent != nil; parent = parent.ParentMessage() {
		name := parent.GetName()
		if len(name) == 0 {
			break
		}
		names = append(names, name)
	}
	names = helper.ReverseStrings(names)
	return names
}
