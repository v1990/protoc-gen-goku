package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type EnumOptions struct {
	common

	pb *descriptorpb.EnumOptions

	uninterpreted_option []*UninterpretedOption
}

func newEnumOptions(desc *descriptorpb.EnumOptions) *EnumOptions {
	t := new(EnumOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *EnumOptions) GetAllowAlias() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetAllowAlias()

}

func (t *EnumOptions) AllowAlias() bool {
	return t.GetAllowAlias()
}

func (t *EnumOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *EnumOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *EnumOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *EnumOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *EnumOptions) PbDescriptor() *descriptorpb.EnumOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumOptions) EnumOptions() *descriptorpb.EnumOptions {
	return t.PbDescriptor()
}

func (t *EnumOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *EnumOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *EnumOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *EnumOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *EnumOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *EnumOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
