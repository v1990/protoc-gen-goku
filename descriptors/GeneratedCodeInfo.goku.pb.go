package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type GeneratedCodeInfo struct {
	common

	pb *descriptorpb.GeneratedCodeInfo

	annotation []*GeneratedCodeInfo_Annotation
}

func newGeneratedCodeInfo(desc *descriptorpb.GeneratedCodeInfo) *GeneratedCodeInfo {
	t := new(GeneratedCodeInfo)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *GeneratedCodeInfo) GetAnnotation() (ret []*GeneratedCodeInfo_Annotation) {
	if t.Empty() {
		return
	}

	if t.annotation != nil {
		return t.annotation
	}

	t.annotation = make([]*GeneratedCodeInfo_Annotation, len(t.pb.GetAnnotation()))

	for i, item := range t.pb.GetAnnotation() {
		elem := newGeneratedCodeInfo_Annotation(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.annotation[i] = elem
	}

	return t.annotation

}

func (t *GeneratedCodeInfo) Annotation() []*GeneratedCodeInfo_Annotation {
	return t.GetAnnotation()
}

func (t *GeneratedCodeInfo) PbDescriptor() *descriptorpb.GeneratedCodeInfo {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *GeneratedCodeInfo) GeneratedCodeInfo() *descriptorpb.GeneratedCodeInfo {
	return t.PbDescriptor()
}

func (t *GeneratedCodeInfo) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *GeneratedCodeInfo) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *GeneratedCodeInfo) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *GeneratedCodeInfo) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *GeneratedCodeInfo) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *GeneratedCodeInfo) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *GeneratedCodeInfo) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
