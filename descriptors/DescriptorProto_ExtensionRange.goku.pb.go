package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type DescriptorProto_ExtensionRange struct {
	common

	pb *descriptorpb.DescriptorProto_ExtensionRange

	options *ExtensionRangeOptions
}

func newDescriptorProto_ExtensionRange(desc *descriptorpb.DescriptorProto_ExtensionRange) *DescriptorProto_ExtensionRange {
	t := new(DescriptorProto_ExtensionRange)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *DescriptorProto_ExtensionRange) GetStart() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetStart()

}

func (t *DescriptorProto_ExtensionRange) Start() int32 {
	return t.GetStart()
}

func (t *DescriptorProto_ExtensionRange) GetEnd() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetEnd()

}

func (t *DescriptorProto_ExtensionRange) End() int32 {
	return t.GetEnd()
}

func (t *DescriptorProto_ExtensionRange) GetOptions() (ret *ExtensionRangeOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newExtensionRangeOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *DescriptorProto_ExtensionRange) Options() *ExtensionRangeOptions {
	return t.GetOptions()
}

func (t *DescriptorProto_ExtensionRange) PbDescriptor() *descriptorpb.DescriptorProto_ExtensionRange {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *DescriptorProto_ExtensionRange) DescriptorProto_ExtensionRange() *descriptorpb.DescriptorProto_ExtensionRange {
	return t.PbDescriptor()
}

func (t *DescriptorProto_ExtensionRange) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *DescriptorProto_ExtensionRange) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *DescriptorProto_ExtensionRange) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *DescriptorProto_ExtensionRange) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *DescriptorProto_ExtensionRange) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *DescriptorProto_ExtensionRange) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *DescriptorProto_ExtensionRange) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
