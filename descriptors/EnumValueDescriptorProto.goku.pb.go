package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type EnumValueDescriptorProto struct {
	common

	pb *descriptorpb.EnumValueDescriptorProto

	options *EnumValueOptions
}

func newEnumValueDescriptorProto(desc *descriptorpb.EnumValueDescriptorProto) *EnumValueDescriptorProto {
	t := new(EnumValueDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *EnumValueDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *EnumValueDescriptorProto) Name() string {
	return t.GetName()
}

func (t *EnumValueDescriptorProto) GetNumber() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetNumber()

}

func (t *EnumValueDescriptorProto) Number() int32 {
	return t.GetNumber()
}

func (t *EnumValueDescriptorProto) GetOptions() (ret *EnumValueOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newEnumValueOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *EnumValueDescriptorProto) Options() *EnumValueOptions {
	return t.GetOptions()
}

func (t *EnumValueDescriptorProto) PbDescriptor() *descriptorpb.EnumValueDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumValueDescriptorProto) EnumValueDescriptorProto() *descriptorpb.EnumValueDescriptorProto {
	return t.PbDescriptor()
}

func (t *EnumValueDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *EnumValueDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *EnumValueDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *EnumValueDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *EnumValueDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *EnumValueDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumValueDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
