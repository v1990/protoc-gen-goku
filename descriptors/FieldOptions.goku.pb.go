package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FieldOptions struct {
	common

	pb *descriptorpb.FieldOptions

	uninterpreted_option []*UninterpretedOption
}

func newFieldOptions(desc *descriptorpb.FieldOptions) *FieldOptions {
	t := new(FieldOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *FieldOptions) GetCtype() (ret FieldOptions_CType) {
	if t.Empty() {
		return
	}

	return t.pb.GetCtype()

}

func (t *FieldOptions) Ctype() FieldOptions_CType {
	return t.GetCtype()
}

func (t *FieldOptions) GetPacked() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetPacked()

}

func (t *FieldOptions) Packed() bool {
	return t.GetPacked()
}

func (t *FieldOptions) GetJstype() (ret FieldOptions_JSType) {
	if t.Empty() {
		return
	}

	return t.pb.GetJstype()

}

func (t *FieldOptions) Jstype() FieldOptions_JSType {
	return t.GetJstype()
}

func (t *FieldOptions) GetLazy() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetLazy()

}

func (t *FieldOptions) Lazy() bool {
	return t.GetLazy()
}

func (t *FieldOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *FieldOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *FieldOptions) GetWeak() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetWeak()

}

func (t *FieldOptions) Weak() bool {
	return t.GetWeak()
}

func (t *FieldOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *FieldOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *FieldOptions) PbDescriptor() *descriptorpb.FieldOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *FieldOptions) FieldOptions() *descriptorpb.FieldOptions {
	return t.PbDescriptor()
}

func (t *FieldOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *FieldOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *FieldOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *FieldOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *FieldOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *FieldOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *FieldOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
