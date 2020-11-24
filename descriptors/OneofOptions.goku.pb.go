package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type OneofOptions struct {
	common

	pb *descriptorpb.OneofOptions

	uninterpreted_option []*UninterpretedOption
}

func newOneofOptions(desc *descriptorpb.OneofOptions) *OneofOptions {
	t := new(OneofOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *OneofOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *OneofOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *OneofOptions) PbDescriptor() *descriptorpb.OneofOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *OneofOptions) OneofOptions() *descriptorpb.OneofOptions {
	return t.PbDescriptor()
}

func (t *OneofOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *OneofOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *OneofOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *OneofOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *OneofOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *OneofOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *OneofOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
