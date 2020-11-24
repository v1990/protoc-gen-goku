package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type ExtensionRangeOptions struct {
	common

	pb *descriptorpb.ExtensionRangeOptions

	uninterpreted_option []*UninterpretedOption
}

func newExtensionRangeOptions(desc *descriptorpb.ExtensionRangeOptions) *ExtensionRangeOptions {
	t := new(ExtensionRangeOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *ExtensionRangeOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *ExtensionRangeOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *ExtensionRangeOptions) PbDescriptor() *descriptorpb.ExtensionRangeOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *ExtensionRangeOptions) ExtensionRangeOptions() *descriptorpb.ExtensionRangeOptions {
	return t.PbDescriptor()
}

func (t *ExtensionRangeOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *ExtensionRangeOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *ExtensionRangeOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *ExtensionRangeOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *ExtensionRangeOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *ExtensionRangeOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *ExtensionRangeOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
