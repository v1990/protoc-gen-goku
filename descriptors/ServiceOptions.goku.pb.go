// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

//   See: descriptorpb.ServiceOptions
type ServiceOptions struct {
	common

	pb *descriptorpb.ServiceOptions

	uninterpreted_option []*UninterpretedOption
}

func newServiceOptions(desc *descriptorpb.ServiceOptions) *ServiceOptions {
	t := new(ServiceOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

// Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
// framework.  We apologize for hoarding these numbers to ourselves, but
// we were already using them long before we decided to release Protocol
// Buffers.  // Is this service deprecated?
// Depending on the target platform, this can emit Deprecated annotations
// for the service, or it will be completely ignored; in the very least,
// this is a formalization for deprecating services.
//   See descriptorpb.ServiceOptions ServiceOptions.Deprecated
//   SourceCodeInfo.Location.Path: [4 16 2 0]
//   proto info: {"name":"deprecated","number":33,"label":"LABEL_OPTIONAL","type":"TYPE_BOOL","defaultValue":"false","jsonName":"deprecated"}
func (t *ServiceOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *ServiceOptions) Deprecated() bool {
	return t.GetDeprecated()
}

// The parser stores options it doesn't recognize here. See above.
//   See descriptorpb.ServiceOptions ServiceOptions.UninterpretedOption
//   SourceCodeInfo.Location.Path: [4 16 2 1]
//   proto info: {"name":"uninterpreted_option","number":999,"label":"LABEL_REPEATED","type":"TYPE_MESSAGE","typeName":".google.protobuf.UninterpretedOption","jsonName":"uninterpretedOption"}
func (t *ServiceOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
	if t.Empty() {
		return
	}

	if t.uninterpreted_option != nil {
		return t.uninterpreted_option
	}

	t.uninterpreted_option = make([]*UninterpretedOption, len(t.pb.GetUninterpretedOption()))

	for i, item := range t.pb.GetUninterpretedOption() {
		elem := newUninterpretedOption(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.uninterpreted_option[i] = elem
	}

	return t.uninterpreted_option

}

func (t *ServiceOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *ServiceOptions) PbDescriptor() *descriptorpb.ServiceOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *ServiceOptions) ServiceOptions() *descriptorpb.ServiceOptions {
	return t.PbDescriptor()
}

func (t *ServiceOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

// implement DescriptorCommon.Empty()
func (t *ServiceOptions) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *ServiceOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *ServiceOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *ServiceOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *ServiceOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}