// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// Describes a method of a service.
//   See: descriptorpb.MethodDescriptorProto
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

//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.Name
//   SourceCodeInfo.Location.Path: [4 9 2 0]
//   proto info: {"name":"name","number":1,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"name"}
func (t *MethodDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *MethodDescriptorProto) Name() string {
	return t.GetName()
}

// Input and output type names.  These are resolved in the same way as
// FieldDescriptorProto.type_name, but must refer to a message type.
//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.InputType
//   SourceCodeInfo.Location.Path: [4 9 2 1]
//   proto info: {"name":"input_type","number":2,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"inputType"}
func (t *MethodDescriptorProto) GetInputType() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetInputType()

}

func (t *MethodDescriptorProto) InputType() string {
	return t.GetInputType()
}

//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.OutputType
//   SourceCodeInfo.Location.Path: [4 9 2 2]
//   proto info: {"name":"output_type","number":3,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"outputType"}
func (t *MethodDescriptorProto) GetOutputType() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetOutputType()

}

func (t *MethodDescriptorProto) OutputType() string {
	return t.GetOutputType()
}

//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.Options
//   SourceCodeInfo.Location.Path: [4 9 2 3]
//   proto info: {"name":"options","number":4,"label":"LABEL_OPTIONAL","type":"TYPE_MESSAGE","typeName":".google.protobuf.MethodOptions","jsonName":"options"}
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

// Identifies if client streams multiple client messages
//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.ClientStreaming
//   SourceCodeInfo.Location.Path: [4 9 2 4]
//   proto info: {"name":"client_streaming","number":5,"label":"LABEL_OPTIONAL","type":"TYPE_BOOL","defaultValue":"false","jsonName":"clientStreaming"}
func (t *MethodDescriptorProto) GetClientStreaming() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetClientStreaming()

}

func (t *MethodDescriptorProto) ClientStreaming() bool {
	return t.GetClientStreaming()
}

// Identifies if server streams multiple server messages
//   See descriptorpb.MethodDescriptorProto MethodDescriptorProto.ServerStreaming
//   SourceCodeInfo.Location.Path: [4 9 2 5]
//   proto info: {"name":"server_streaming","number":6,"label":"LABEL_OPTIONAL","type":"TYPE_BOOL","defaultValue":"false","jsonName":"serverStreaming"}
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

// implement DescriptorCommon.Empty()
func (t *MethodDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *MethodDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *MethodDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *MethodDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *MethodDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}