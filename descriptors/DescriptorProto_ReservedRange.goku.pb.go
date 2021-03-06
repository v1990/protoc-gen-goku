package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type DescriptorProto_ReservedRange struct {
	common

	pb *descriptorpb.DescriptorProto_ReservedRange
}

func newDescriptorProto_ReservedRange(desc *descriptorpb.DescriptorProto_ReservedRange) *DescriptorProto_ReservedRange {
	t := new(DescriptorProto_ReservedRange)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *DescriptorProto_ReservedRange) GetStart() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetStart()

}

func (t *DescriptorProto_ReservedRange) Start() int32 {
	return t.GetStart()
}

func (t *DescriptorProto_ReservedRange) GetEnd() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetEnd()

}

func (t *DescriptorProto_ReservedRange) End() int32 {
	return t.GetEnd()
}

func (t *DescriptorProto_ReservedRange) PbDescriptor() *descriptorpb.DescriptorProto_ReservedRange {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *DescriptorProto_ReservedRange) DescriptorProto_ReservedRange() *descriptorpb.DescriptorProto_ReservedRange {
	return t.PbDescriptor()
}

func (t *DescriptorProto_ReservedRange) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *DescriptorProto_ReservedRange) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *DescriptorProto_ReservedRange) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *DescriptorProto_ReservedRange) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *DescriptorProto_ReservedRange) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *DescriptorProto_ReservedRange) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *DescriptorProto_ReservedRange) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
