package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type SourceCodeInfo_Location struct {
	common

	pb *descriptorpb.SourceCodeInfo_Location
}

func newSourceCodeInfo_Location(desc *descriptorpb.SourceCodeInfo_Location) *SourceCodeInfo_Location {
	t := new(SourceCodeInfo_Location)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *SourceCodeInfo_Location) GetPath() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetPath()

}

func (t *SourceCodeInfo_Location) Path() []int32 {
	return t.GetPath()
}

func (t *SourceCodeInfo_Location) GetSpan() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetSpan()

}

func (t *SourceCodeInfo_Location) Span() []int32 {
	return t.GetSpan()
}

func (t *SourceCodeInfo_Location) GetLeadingComments() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetLeadingComments()

}

func (t *SourceCodeInfo_Location) LeadingComments() string {
	return t.GetLeadingComments()
}

func (t *SourceCodeInfo_Location) GetTrailingComments() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetTrailingComments()

}

func (t *SourceCodeInfo_Location) TrailingComments() string {
	return t.GetTrailingComments()
}

func (t *SourceCodeInfo_Location) GetLeadingDetachedComments() (ret []string) {
	if t.Empty() {
		return
	}

	return t.pb.GetLeadingDetachedComments()

}

func (t *SourceCodeInfo_Location) LeadingDetachedComments() []string {
	return t.GetLeadingDetachedComments()
}

func (t *SourceCodeInfo_Location) PbDescriptor() *descriptorpb.SourceCodeInfo_Location {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *SourceCodeInfo_Location) SourceCodeInfo_Location() *descriptorpb.SourceCodeInfo_Location {
	return t.PbDescriptor()
}

func (t *SourceCodeInfo_Location) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *SourceCodeInfo_Location) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *SourceCodeInfo_Location) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *SourceCodeInfo_Location) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *SourceCodeInfo_Location) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *SourceCodeInfo_Location) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *SourceCodeInfo_Location) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
