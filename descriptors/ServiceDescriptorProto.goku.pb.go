package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type ServiceDescriptorProto struct {
	common

	pb *descriptorpb.ServiceDescriptorProto

	method []*MethodDescriptorProto

	options *ServiceOptions
}

func newServiceDescriptorProto(desc *descriptorpb.ServiceDescriptorProto) *ServiceDescriptorProto {
	t := new(ServiceDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *ServiceDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *ServiceDescriptorProto) Name() string {
	return t.GetName()
}

func (t *ServiceDescriptorProto) GetMethod() (ret []*MethodDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.method != nil {
		return t.method
	}

	t.method = make([]*MethodDescriptorProto, len(t.pb.GetMethod()))

	for i, item := range t.pb.GetMethod() {
		elem := newMethodDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.method[i] = elem
	}

	return t.method

}

func (t *ServiceDescriptorProto) Method() []*MethodDescriptorProto {
	return t.GetMethod()
}

func (t *ServiceDescriptorProto) GetOptions() (ret *ServiceOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newServiceOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *ServiceDescriptorProto) Options() *ServiceOptions {
	return t.GetOptions()
}

func (t *ServiceDescriptorProto) PbDescriptor() *descriptorpb.ServiceDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *ServiceDescriptorProto) ServiceDescriptorProto() *descriptorpb.ServiceDescriptorProto {
	return t.PbDescriptor()
}

func (t *ServiceDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *ServiceDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *ServiceDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *ServiceDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *ServiceDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *ServiceDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *ServiceDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
