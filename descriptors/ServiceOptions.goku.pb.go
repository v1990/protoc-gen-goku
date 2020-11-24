package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

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

func (t *ServiceOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *ServiceOptions) Deprecated() bool {
	return t.GetDeprecated()
}

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

func (t *ServiceOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *ServiceOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *ServiceOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *ServiceOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *ServiceOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *ServiceOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
