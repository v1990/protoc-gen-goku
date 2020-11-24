package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MethodOptions struct {
	common

	pb *descriptorpb.MethodOptions

	uninterpreted_option []*UninterpretedOption
}

func newMethodOptions(desc *descriptorpb.MethodOptions) *MethodOptions {
	t := new(MethodOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *MethodOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *MethodOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *MethodOptions) GetIdempotencyLevel() (ret MethodOptions_IdempotencyLevel) {
	if t.Empty() {
		return
	}

	return t.pb.GetIdempotencyLevel()

}

func (t *MethodOptions) IdempotencyLevel() MethodOptions_IdempotencyLevel {
	return t.GetIdempotencyLevel()
}

func (t *MethodOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *MethodOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *MethodOptions) PbDescriptor() *descriptorpb.MethodOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *MethodOptions) MethodOptions() *descriptorpb.MethodOptions {
	return t.PbDescriptor()
}

func (t *MethodOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *MethodOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *MethodOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *MethodOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *MethodOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *MethodOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *MethodOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
