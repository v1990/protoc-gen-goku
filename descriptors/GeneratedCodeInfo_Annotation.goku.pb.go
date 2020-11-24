package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type GeneratedCodeInfo_Annotation struct {
	common

	pb *descriptorpb.GeneratedCodeInfo_Annotation
}

func newGeneratedCodeInfo_Annotation(desc *descriptorpb.GeneratedCodeInfo_Annotation) *GeneratedCodeInfo_Annotation {
	t := new(GeneratedCodeInfo_Annotation)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *GeneratedCodeInfo_Annotation) GetPath() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetPath()

}

func (t *GeneratedCodeInfo_Annotation) Path() []int32 {
	return t.GetPath()
}

func (t *GeneratedCodeInfo_Annotation) GetSourceFile() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetSourceFile()

}

func (t *GeneratedCodeInfo_Annotation) SourceFile() string {
	return t.GetSourceFile()
}

func (t *GeneratedCodeInfo_Annotation) GetBegin() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetBegin()

}

func (t *GeneratedCodeInfo_Annotation) Begin() int32 {
	return t.GetBegin()
}

func (t *GeneratedCodeInfo_Annotation) GetEnd() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetEnd()

}

func (t *GeneratedCodeInfo_Annotation) End() int32 {
	return t.GetEnd()
}

func (t *GeneratedCodeInfo_Annotation) PbDescriptor() *descriptorpb.GeneratedCodeInfo_Annotation {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *GeneratedCodeInfo_Annotation) GeneratedCodeInfo_Annotation() *descriptorpb.GeneratedCodeInfo_Annotation {
	return t.PbDescriptor()
}

func (t *GeneratedCodeInfo_Annotation) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *GeneratedCodeInfo_Annotation) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *GeneratedCodeInfo_Annotation) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *GeneratedCodeInfo_Annotation) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *GeneratedCodeInfo_Annotation) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *GeneratedCodeInfo_Annotation) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *GeneratedCodeInfo_Annotation) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
