package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type SourceCodeInfo struct {
	common

	pb *descriptorpb.SourceCodeInfo

	location []*SourceCodeInfo_Location
}

func newSourceCodeInfo(desc *descriptorpb.SourceCodeInfo) *SourceCodeInfo {
	t := new(SourceCodeInfo)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *SourceCodeInfo) GetLocation() (ret []*SourceCodeInfo_Location) {
	if t.Empty() {
		return
	}

	if t.location != nil {
		return t.location
	}

	t.location = make([]*SourceCodeInfo_Location, len(t.pb.GetLocation()))

	for i, item := range t.pb.GetLocation() {
		elem := newSourceCodeInfo_Location(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.location[i] = elem
	}

	return t.location

}

func (t *SourceCodeInfo) Location() []*SourceCodeInfo_Location {
	return t.GetLocation()
}

func (t *SourceCodeInfo) PbDescriptor() *descriptorpb.SourceCodeInfo {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *SourceCodeInfo) SourceCodeInfo() *descriptorpb.SourceCodeInfo {
	return t.PbDescriptor()
}

func (t *SourceCodeInfo) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *SourceCodeInfo) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *SourceCodeInfo) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *SourceCodeInfo) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *SourceCodeInfo) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *SourceCodeInfo) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *SourceCodeInfo) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
