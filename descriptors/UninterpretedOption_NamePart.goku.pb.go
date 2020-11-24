package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type UninterpretedOption_NamePart struct {
	common

	pb *descriptorpb.UninterpretedOption_NamePart
}

func newUninterpretedOption_NamePart(desc *descriptorpb.UninterpretedOption_NamePart) *UninterpretedOption_NamePart {
	t := new(UninterpretedOption_NamePart)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *UninterpretedOption_NamePart) GetNamePart() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetNamePart()

}

func (t *UninterpretedOption_NamePart) NamePart() string {
	return t.GetNamePart()
}

func (t *UninterpretedOption_NamePart) GetIsExtension() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetIsExtension()

}

func (t *UninterpretedOption_NamePart) IsExtension() bool {
	return t.GetIsExtension()
}

func (t *UninterpretedOption_NamePart) PbDescriptor() *descriptorpb.UninterpretedOption_NamePart {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *UninterpretedOption_NamePart) UninterpretedOption_NamePart() *descriptorpb.UninterpretedOption_NamePart {
	return t.PbDescriptor()
}

func (t *UninterpretedOption_NamePart) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *UninterpretedOption_NamePart) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *UninterpretedOption_NamePart) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *UninterpretedOption_NamePart) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *UninterpretedOption_NamePart) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *UninterpretedOption_NamePart) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *UninterpretedOption_NamePart) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
