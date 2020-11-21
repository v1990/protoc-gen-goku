// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// Describes an enum type.
//   See: descriptorpb.EnumDescriptorProto
type EnumDescriptorProto struct {
	common

	pb *descriptorpb.EnumDescriptorProto

	value []*EnumValueDescriptorProto

	options *EnumOptions

	reserved_range []*EnumDescriptorProto_EnumReservedRange

	pbTypeInfo *PbTypeInfo
}

func newEnumDescriptorProto(desc *descriptorpb.EnumDescriptorProto) *EnumDescriptorProto {
	t := new(EnumDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

//   See descriptorpb.EnumDescriptorProto EnumDescriptorProto.Name
//   SourceCodeInfo.Location.Path: [4 6 2 0]
//   proto info: {"name":"name","number":1,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"name"}
func (t *EnumDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *EnumDescriptorProto) Name() string {
	return t.GetName()
}

//   See descriptorpb.EnumDescriptorProto EnumDescriptorProto.Value
//   SourceCodeInfo.Location.Path: [4 6 2 1]
//   proto info: {"name":"value","number":2,"label":"LABEL_REPEATED","type":"TYPE_MESSAGE","typeName":".google.protobuf.EnumValueDescriptorProto","jsonName":"value"}
func (t *EnumDescriptorProto) GetValue() (ret []*EnumValueDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.value != nil {
		return t.value
	}

	t.value = make([]*EnumValueDescriptorProto, len(t.pb.GetValue()))

	for i, item := range t.pb.GetValue() {
		elem := newEnumValueDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.value[i] = elem
	}

	return t.value

}

func (t *EnumDescriptorProto) Value() []*EnumValueDescriptorProto {
	return t.GetValue()
}

//   See descriptorpb.EnumDescriptorProto EnumDescriptorProto.Options
//   SourceCodeInfo.Location.Path: [4 6 2 2]
//   proto info: {"name":"options","number":3,"label":"LABEL_OPTIONAL","type":"TYPE_MESSAGE","typeName":".google.protobuf.EnumOptions","jsonName":"options"}
func (t *EnumDescriptorProto) GetOptions() (ret *EnumOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newEnumOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *EnumDescriptorProto) Options() *EnumOptions {
	return t.GetOptions()
}

// Range of reserved numeric values. Reserved numeric values may not be used
// by enum values in the same enum declaration. Reserved ranges may not
// overlap.
//   See descriptorpb.EnumDescriptorProto EnumDescriptorProto.ReservedRange
//   SourceCodeInfo.Location.Path: [4 6 2 3]
//   proto info: {"name":"reserved_range","number":4,"label":"LABEL_REPEATED","type":"TYPE_MESSAGE","typeName":".google.protobuf.EnumDescriptorProto.EnumReservedRange","jsonName":"reservedRange"}
func (t *EnumDescriptorProto) GetReservedRange() (ret []*EnumDescriptorProto_EnumReservedRange) {
	if t.Empty() {
		return
	}

	if t.reserved_range != nil {
		return t.reserved_range
	}

	t.reserved_range = make([]*EnumDescriptorProto_EnumReservedRange, len(t.pb.GetReservedRange()))

	for i, item := range t.pb.GetReservedRange() {
		elem := newEnumDescriptorProto_EnumReservedRange(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.reserved_range[i] = elem
	}

	return t.reserved_range

}

func (t *EnumDescriptorProto) ReservedRange() []*EnumDescriptorProto_EnumReservedRange {
	return t.GetReservedRange()
}

// Reserved enum value names, which may not be reused. A given name may only
// be reserved once.
//   See descriptorpb.EnumDescriptorProto EnumDescriptorProto.ReservedName
//   SourceCodeInfo.Location.Path: [4 6 2 4]
//   proto info: {"name":"reserved_name","number":5,"label":"LABEL_REPEATED","type":"TYPE_STRING","jsonName":"reservedName"}
func (t *EnumDescriptorProto) GetReservedName() (ret []string) {
	if t.Empty() {
		return
	}

	return t.pb.GetReservedName()

}

func (t *EnumDescriptorProto) ReservedName() []string {
	return t.GetReservedName()
}

func (t *EnumDescriptorProto) PbDescriptor() *descriptorpb.EnumDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumDescriptorProto) EnumDescriptorProto() *descriptorpb.EnumDescriptorProto {
	return t.PbDescriptor()
}

func (t *EnumDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

// implement DescriptorCommon.Empty()
func (t *EnumDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *EnumDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *EnumDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *EnumDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *EnumDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

// implement Nestable.IsNested()
func (t *EnumDescriptorProto) IsNested() bool {
	if t.Empty() {
		return false
	}
	return t.isNested()
}

// implement Nestable.ParentMessage()
func (t *EnumDescriptorProto) ParentMessage() *DescriptorProto {
	if t.Empty() {
		return nil
	}
	return t.parentMessage()
}

// implemented ProtoType.ProtoType()
func (t *EnumDescriptorProto) ProtoType() *PbTypeInfo {
	if t.Empty() {
		return nil
	}
	if t.pbTypeInfo == nil {
		t.pbTypeInfo = &PbTypeInfo{d: t, names: nestedTypeNames(t)}
	}
	return t.pbTypeInfo
}
