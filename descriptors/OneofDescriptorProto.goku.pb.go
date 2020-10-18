// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// Describes a oneof.
//   See: descriptorpb.OneofDescriptorProto
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

//   See descriptorpb.OneofDescriptorProto OneofDescriptorProto.Name
//   SourceCodeInfo.Location.Path: [4 5 2 0]
//   proto info: {"name":"name","number":1,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"name"}
func (t *OneofDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *OneofDescriptorProto) Name() string {
	return t.GetName()
}

//   See descriptorpb.OneofDescriptorProto OneofDescriptorProto.Options
//   SourceCodeInfo.Location.Path: [4 5 2 1]
//   proto info: {"name":"options","number":2,"label":"LABEL_OPTIONAL","type":"TYPE_MESSAGE","typeName":".google.protobuf.OneofOptions","jsonName":"options"}
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

// implement DescriptorCommon.Empty()
func (t *OneofDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *OneofDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *OneofDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *OneofDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *OneofDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}