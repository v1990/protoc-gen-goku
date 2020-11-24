package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MessageOptions struct {
	common

	pb *descriptorpb.MessageOptions

	uninterpreted_option []*UninterpretedOption
}

func newMessageOptions(desc *descriptorpb.MessageOptions) *MessageOptions {
	t := new(MessageOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *MessageOptions) GetMessageSetWireFormat() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetMessageSetWireFormat()

}

func (t *MessageOptions) MessageSetWireFormat() bool {
	return t.GetMessageSetWireFormat()
}

func (t *MessageOptions) GetNoStandardDescriptorAccessor() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetNoStandardDescriptorAccessor()

}

func (t *MessageOptions) NoStandardDescriptorAccessor() bool {
	return t.GetNoStandardDescriptorAccessor()
}

func (t *MessageOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *MessageOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *MessageOptions) GetMapEntry() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetMapEntry()

}

func (t *MessageOptions) MapEntry() bool {
	return t.GetMapEntry()
}

func (t *MessageOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
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

func (t *MessageOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *MessageOptions) PbDescriptor() *descriptorpb.MessageOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *MessageOptions) MessageOptions() *descriptorpb.MessageOptions {
	return t.PbDescriptor()
}

func (t *MessageOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *MessageOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *MessageOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *MessageOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *MessageOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *MessageOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *MessageOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
