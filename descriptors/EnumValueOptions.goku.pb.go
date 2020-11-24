package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type EnumValueOptions struct {
	common

	pb *descriptorpb.EnumValueOptions

	uninterpreted_option []*UninterpretedOption
}

func newEnumValueOptions(desc *descriptorpb.EnumValueOptions) *EnumValueOptions {
	t := new(EnumValueOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *EnumValueOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *EnumValueOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *EnumValueOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *EnumValueOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *EnumValueOptions) PbDescriptor() *descriptorpb.EnumValueOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumValueOptions) EnumValueOptions() *descriptorpb.EnumValueOptions {
	return t.PbDescriptor()
}

func (t *EnumValueOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *EnumValueOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *EnumValueOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *EnumValueOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *EnumValueOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *EnumValueOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumValueOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
