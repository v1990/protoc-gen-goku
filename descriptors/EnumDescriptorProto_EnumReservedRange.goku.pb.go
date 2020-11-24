package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type EnumDescriptorProto_EnumReservedRange struct {
	common

	pb *descriptorpb.EnumDescriptorProto_EnumReservedRange
}

func newEnumDescriptorProto_EnumReservedRange(desc *descriptorpb.EnumDescriptorProto_EnumReservedRange) *EnumDescriptorProto_EnumReservedRange {
	t := new(EnumDescriptorProto_EnumReservedRange)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *EnumDescriptorProto_EnumReservedRange) GetStart() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetStart()

}

func (t *EnumDescriptorProto_EnumReservedRange) Start() int32 {
	return t.GetStart()
}

func (t *EnumDescriptorProto_EnumReservedRange) GetEnd() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetEnd()

}

func (t *EnumDescriptorProto_EnumReservedRange) End() int32 {
	return t.GetEnd()
}

func (t *EnumDescriptorProto_EnumReservedRange) PbDescriptor() *descriptorpb.EnumDescriptorProto_EnumReservedRange {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumDescriptorProto_EnumReservedRange) EnumDescriptorProto_EnumReservedRange() *descriptorpb.EnumDescriptorProto_EnumReservedRange {
	return t.PbDescriptor()
}

func (t *EnumDescriptorProto_EnumReservedRange) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *EnumDescriptorProto_EnumReservedRange) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *EnumDescriptorProto_EnumReservedRange) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *EnumDescriptorProto_EnumReservedRange) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *EnumDescriptorProto_EnumReservedRange) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *EnumDescriptorProto_EnumReservedRange) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumDescriptorProto_EnumReservedRange) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
