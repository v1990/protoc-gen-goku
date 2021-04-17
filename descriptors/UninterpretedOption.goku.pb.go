// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// A message representing a option the parser does not recognize. This only
// appears in options protos created by the compiler::Parser class.
// DescriptorPool resolves these when building Descriptor objects. Therefore,
// options protos in descriptor objects (e.g. returned by Descriptor::options(),
// or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
// in them.
//   See: descriptorpb.UninterpretedOption
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

//   See descriptorpb.UninterpretedOption UninterpretedOption.Name
//   SourceCodeInfo.Location.Path: [4 18 2 0]
//   proto info: {}
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

// The value of the uninterpreted option, in whatever type the tokenizer
// identified it as during parsing. Exactly one of these should be set.
//   See descriptorpb.UninterpretedOption UninterpretedOption.IdentifierValue
//   SourceCodeInfo.Location.Path: [4 18 2 1]
//   proto info: {}
func (t *UninterpretedOption) GetIdentifierValue() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetIdentifierValue()

}

func (t *UninterpretedOption) IdentifierValue() string {
	return t.GetIdentifierValue()
}

//   See descriptorpb.UninterpretedOption UninterpretedOption.PositiveIntValue
//   SourceCodeInfo.Location.Path: [4 18 2 2]
//   proto info: {}
func (t *UninterpretedOption) GetPositiveIntValue() (ret uint64) {
	if t.Empty() {
		return
	}

	return t.pb.GetPositiveIntValue()

}

func (t *UninterpretedOption) PositiveIntValue() uint64 {
	return t.GetPositiveIntValue()
}

//   See descriptorpb.UninterpretedOption UninterpretedOption.NegativeIntValue
//   SourceCodeInfo.Location.Path: [4 18 2 3]
//   proto info: {}
func (t *UninterpretedOption) GetNegativeIntValue() (ret int64) {
	if t.Empty() {
		return
	}

	return t.pb.GetNegativeIntValue()

}

func (t *UninterpretedOption) NegativeIntValue() int64 {
	return t.GetNegativeIntValue()
}

//   See descriptorpb.UninterpretedOption UninterpretedOption.DoubleValue
//   SourceCodeInfo.Location.Path: [4 18 2 4]
//   proto info: {}
func (t *UninterpretedOption) GetDoubleValue() (ret float64) {
	if t.Empty() {
		return
	}

	return t.pb.GetDoubleValue()

}

func (t *UninterpretedOption) DoubleValue() float64 {
	return t.GetDoubleValue()
}

//   See descriptorpb.UninterpretedOption UninterpretedOption.StringValue
//   SourceCodeInfo.Location.Path: [4 18 2 5]
//   proto info: {}
func (t *UninterpretedOption) GetStringValue() (ret []byte) {
	if t.Empty() {
		return
	}

	return t.pb.GetStringValue()

}

func (t *UninterpretedOption) StringValue() []byte {
	return t.GetStringValue()
}

//   See descriptorpb.UninterpretedOption UninterpretedOption.AggregateValue
//   SourceCodeInfo.Location.Path: [4 18 2 6]
//   proto info: {}
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

/*
func (t *UninterpretedOption) MarshalJSON() (b []byte,err error) {
    if t.Empty() {
        return
    }
    buf := bytes.NewBuffer(nil)
    err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
    return buf.Bytes(), err
}
*/

// ExportFields returns can export fields
func (t *UninterpretedOption) ExportFields() map[string]interface{} {
	return map[string]interface{}{
		"Name":             t.Name(),
		"IdentifierValue":  t.IdentifierValue(),
		"PositiveIntValue": t.PositiveIntValue(),
		"NegativeIntValue": t.NegativeIntValue(),
		"DoubleValue":      t.DoubleValue(),
		"StringValue":      t.StringValue(),
		"AggregateValue":   t.AggregateValue(),
	}
}

// implement DescriptorCommon.Empty()
func (t *UninterpretedOption) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *UninterpretedOption) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *UninterpretedOption) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *UninterpretedOption) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
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
