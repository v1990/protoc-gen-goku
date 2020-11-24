package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

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

func (t *EnumDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *EnumDescriptorProto) Name() string {
	return t.GetName()
}

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

func (t *EnumDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *EnumDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *EnumDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *EnumDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *EnumDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}

func (t *EnumDescriptorProto) IsNested() bool {
	if t.Empty() {
		return false
	}
	return t.isNested()
}

func (t *EnumDescriptorProto) ParentMessage() *DescriptorProto {
	if t.Empty() {
		return nil
	}
	return t.parentMessage()
}

func (t *EnumDescriptorProto) ProtoType() *PbTypeInfo {
	if t.Empty() {
		return nil
	}
	if t.pbTypeInfo == nil {
		t.pbTypeInfo = &PbTypeInfo{d: t, names: nestedTypeNames(t)}
	}
	return t.pbTypeInfo
}
