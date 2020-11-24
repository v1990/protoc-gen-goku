package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type UninterpretedOption struct {
	common

	pb *descriptorpb.UninterpretedOption

	name []*UninterpretedOption_NamePart
}

func newUninterpretedOption(desc *descriptorpb.UninterpretedOption) *UninterpretedOption {
	t := new(UninterpretedOption)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *UninterpretedOption) GetName() (ret []*UninterpretedOption_NamePart) {
	if t.Empty() {
		return
	}

	if t.name != nil {
		return t.name
	}

	t.name = make([]*UninterpretedOption_NamePart, len(t.pb.GetName()))

	for i, item := range t.pb.GetName() {
		elem := newUninterpretedOption_NamePart(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.name[i] = elem
	}

	return t.name

}

func (t *UninterpretedOption) Name() []*UninterpretedOption_NamePart {
	return t.GetName()
}

func (t *UninterpretedOption) GetIdentifierValue() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetIdentifierValue()

}

func (t *UninterpretedOption) IdentifierValue() string {
	return t.GetIdentifierValue()
}

func (t *UninterpretedOption) GetPositiveIntValue() (ret uint64) {
	if t.Empty() {
		return
	}

	return t.pb.GetPositiveIntValue()

}

func (t *UninterpretedOption) PositiveIntValue() uint64 {
	return t.GetPositiveIntValue()
}

func (t *UninterpretedOption) GetNegativeIntValue() (ret int64) {
	if t.Empty() {
		return
	}

	return t.pb.GetNegativeIntValue()

}

func (t *UninterpretedOption) NegativeIntValue() int64 {
	return t.GetNegativeIntValue()
}

func (t *UninterpretedOption) GetDoubleValue() (ret float64) {
	if t.Empty() {
		return
	}

	return t.pb.GetDoubleValue()

}

func (t *UninterpretedOption) DoubleValue() float64 {
	return t.GetDoubleValue()
}

func (t *UninterpretedOption) GetStringValue() (ret []byte) {
	if t.Empty() {
		return
	}

	return t.pb.GetStringValue()

}

func (t *UninterpretedOption) StringValue() []byte {
	return t.GetStringValue()
}

func (t *UninterpretedOption) GetAggregateValue() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetAggregateValue()

}

func (t *UninterpretedOption) AggregateValue() string {
	return t.GetAggregateValue()
}

func (t *UninterpretedOption) PbDescriptor() *descriptorpb.UninterpretedOption {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *UninterpretedOption) UninterpretedOption() *descriptorpb.UninterpretedOption {
	return t.PbDescriptor()
}

func (t *UninterpretedOption) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *UninterpretedOption) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *UninterpretedOption) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *UninterpretedOption) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *UninterpretedOption) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *UninterpretedOption) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *UninterpretedOption) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
