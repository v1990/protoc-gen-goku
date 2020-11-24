package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MethodDescriptorProto struct {
	common

	pb *descriptorpb.MethodDescriptorProto

	options *MethodOptions
}

func newMethodDescriptorProto(desc *descriptorpb.MethodDescriptorProto) *MethodDescriptorProto {
	t := new(MethodDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *MethodDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *MethodDescriptorProto) Name() string {
	return t.GetName()
}

func (t *MethodDescriptorProto) GetInputType() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetInputType()

}

func (t *MethodDescriptorProto) InputType() string {
	return t.GetInputType()
}

func (t *MethodDescriptorProto) GetOutputType() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetOutputType()

}

func (t *MethodDescriptorProto) OutputType() string {
	return t.GetOutputType()
}

func (t *MethodDescriptorProto) GetOptions() (ret *MethodOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newMethodOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *MethodDescriptorProto) Options() *MethodOptions {
	return t.GetOptions()
}

func (t *MethodDescriptorProto) GetClientStreaming() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetClientStreaming()

}

func (t *MethodDescriptorProto) ClientStreaming() bool {
	return t.GetClientStreaming()
}

func (t *MethodDescriptorProto) GetServerStreaming() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetServerStreaming()

}

func (t *MethodDescriptorProto) ServerStreaming() bool {
	return t.GetServerStreaming()
}

func (t *MethodDescriptorProto) PbDescriptor() *descriptorpb.MethodDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *MethodDescriptorProto) MethodDescriptorProto() *descriptorpb.MethodDescriptorProto {
	return t.PbDescriptor()
}

func (t *MethodDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *MethodDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *MethodDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *MethodDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *MethodDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *MethodDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *MethodDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
