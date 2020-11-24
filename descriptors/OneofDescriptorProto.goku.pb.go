package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type OneofDescriptorProto struct {
	common

	pb *descriptorpb.OneofDescriptorProto

	options *OneofOptions
}

func newOneofDescriptorProto(desc *descriptorpb.OneofDescriptorProto) *OneofDescriptorProto {
	t := new(OneofDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *OneofDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *OneofDescriptorProto) Name() string {
	return t.GetName()
}

func (t *OneofDescriptorProto) GetOptions() (ret *OneofOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newOneofOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *OneofDescriptorProto) Options() *OneofOptions {
	return t.GetOptions()
}

func (t *OneofDescriptorProto) PbDescriptor() *descriptorpb.OneofDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *OneofDescriptorProto) OneofDescriptorProto() *descriptorpb.OneofDescriptorProto {
	return t.PbDescriptor()
}

func (t *OneofDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *OneofDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *OneofDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *OneofDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *OneofDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *OneofDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *OneofDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
