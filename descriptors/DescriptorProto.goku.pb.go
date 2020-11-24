package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type DescriptorProto struct {
	common

	pb *descriptorpb.DescriptorProto

	field []*FieldDescriptorProto

	extension []*FieldDescriptorProto

	nested_type []*DescriptorProto

	enum_type []*EnumDescriptorProto

	extension_range []*DescriptorProto_ExtensionRange

	oneof_decl []*OneofDescriptorProto

	options *MessageOptions

	reserved_range []*DescriptorProto_ReservedRange

	pbTypeInfo *PbTypeInfo
}

func newDescriptorProto(desc *descriptorpb.DescriptorProto) *DescriptorProto {
	t := new(DescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *DescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *DescriptorProto) Name() string {
	return t.GetName()
}

func (t *DescriptorProto) GetField() (ret []*FieldDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.field != nil {
		return t.field
	}

	t.field = make([]*FieldDescriptorProto, len(t.pb.GetField()))

	for i, item := range t.pb.GetField() {
		elem := newFieldDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.field[i] = elem
	}

	return t.field

}

func (t *DescriptorProto) Field() []*FieldDescriptorProto {
	return t.GetField()
}

func (t *DescriptorProto) GetExtension() (ret []*FieldDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.extension != nil {
		return t.extension
	}

	t.extension = make([]*FieldDescriptorProto, len(t.pb.GetExtension()))

	for i, item := range t.pb.GetExtension() {
		elem := newFieldDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.extension[i] = elem
	}

	return t.extension

}

func (t *DescriptorProto) Extension() []*FieldDescriptorProto {
	return t.GetExtension()
}

func (t *DescriptorProto) GetNestedType() (ret []*DescriptorProto) {
	if t.Empty() {
		return
	}

	if t.nested_type != nil {
		return t.nested_type
	}

	t.nested_type = make([]*DescriptorProto, len(t.pb.GetNestedType()))

	for i, item := range t.pb.GetNestedType() {
		elem := newDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.nested_type[i] = elem
	}

	return t.nested_type

}

func (t *DescriptorProto) NestedType() []*DescriptorProto {
	return t.GetNestedType()
}

func (t *DescriptorProto) GetEnumType() (ret []*EnumDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.enum_type != nil {
		return t.enum_type
	}

	t.enum_type = make([]*EnumDescriptorProto, len(t.pb.GetEnumType()))

	for i, item := range t.pb.GetEnumType() {
		elem := newEnumDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.enum_type[i] = elem
	}

	return t.enum_type

}

func (t *DescriptorProto) EnumType() []*EnumDescriptorProto {
	return t.GetEnumType()
}

func (t *DescriptorProto) GetExtensionRange() (ret []*DescriptorProto_ExtensionRange) {
	if t.Empty() {
		return
	}

	if t.extension_range != nil {
		return t.extension_range
	}

	t.extension_range = make([]*DescriptorProto_ExtensionRange, len(t.pb.GetExtensionRange()))

	for i, item := range t.pb.GetExtensionRange() {
		elem := newDescriptorProto_ExtensionRange(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.extension_range[i] = elem
	}

	return t.extension_range

}

func (t *DescriptorProto) ExtensionRange() []*DescriptorProto_ExtensionRange {
	return t.GetExtensionRange()
}

func (t *DescriptorProto) GetOneofDecl() (ret []*OneofDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.oneof_decl != nil {
		return t.oneof_decl
	}

	t.oneof_decl = make([]*OneofDescriptorProto, len(t.pb.GetOneofDecl()))

	for i, item := range t.pb.GetOneofDecl() {
		elem := newOneofDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.oneof_decl[i] = elem
	}

	return t.oneof_decl

}

func (t *DescriptorProto) OneofDecl() []*OneofDescriptorProto {
	return t.GetOneofDecl()
}

func (t *DescriptorProto) GetOptions() (ret *MessageOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newMessageOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *DescriptorProto) Options() *MessageOptions {
	return t.GetOptions()
}

func (t *DescriptorProto) GetReservedRange() (ret []*DescriptorProto_ReservedRange) {
	if t.Empty() {
		return
	}

	if t.reserved_range != nil {
		return t.reserved_range
	}

	t.reserved_range = make([]*DescriptorProto_ReservedRange, len(t.pb.GetReservedRange()))

	for i, item := range t.pb.GetReservedRange() {
		elem := newDescriptorProto_ReservedRange(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.reserved_range[i] = elem
	}

	return t.reserved_range

}

func (t *DescriptorProto) ReservedRange() []*DescriptorProto_ReservedRange {
	return t.GetReservedRange()
}

func (t *DescriptorProto) GetReservedName() (ret []string) {
	if t.Empty() {
		return
	}

	return t.pb.GetReservedName()

}

func (t *DescriptorProto) ReservedName() []string {
	return t.GetReservedName()
}

func (t *DescriptorProto) PbDescriptor() *descriptorpb.DescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *DescriptorProto) DescriptorProto() *descriptorpb.DescriptorProto {
	return t.PbDescriptor()
}

func (t *DescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *DescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *DescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *DescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *DescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *DescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *DescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}

func (t *DescriptorProto) IsNested() bool {
	if t.Empty() {
		return false
	}
	return t.isNested()
}

func (t *DescriptorProto) ParentMessage() *DescriptorProto {
	if t.Empty() {
		return nil
	}
	return t.parentMessage()
}

func (t *DescriptorProto) ProtoType() *PbTypeInfo {
	if t.Empty() {
		return nil
	}
	if t.pbTypeInfo == nil {
		t.pbTypeInfo = &PbTypeInfo{d: t, names: nestedTypeNames(t)}
	}
	return t.pbTypeInfo
}
