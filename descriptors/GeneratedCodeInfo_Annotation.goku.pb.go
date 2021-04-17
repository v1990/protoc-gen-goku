// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

//   See: descriptorpb.GeneratedCodeInfo_Annotation
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

// Identifies the element in the original source .proto file. This field
// is formatted the same as SourceCodeInfo.Location.path.
//   See descriptorpb.GeneratedCodeInfo_Annotation GeneratedCodeInfo_Annotation.Path
//   SourceCodeInfo.Location.Path: [4 20 3 0 2 0]
//   proto info: {}
func (t *GeneratedCodeInfo_Annotation) GetPath() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetPath()

}

func (t *GeneratedCodeInfo_Annotation) Path() []int32 {
	return t.GetPath()
}

// Identifies the filesystem path to the original source .proto.
//   See descriptorpb.GeneratedCodeInfo_Annotation GeneratedCodeInfo_Annotation.SourceFile
//   SourceCodeInfo.Location.Path: [4 20 3 0 2 1]
//   proto info: {}
func (t *GeneratedCodeInfo_Annotation) GetSourceFile() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetSourceFile()

}

func (t *GeneratedCodeInfo_Annotation) SourceFile() string {
	return t.GetSourceFile()
}

// Identifies the starting offset in bytes in the generated code
// that relates to the identified object.
//   See descriptorpb.GeneratedCodeInfo_Annotation GeneratedCodeInfo_Annotation.Begin
//   SourceCodeInfo.Location.Path: [4 20 3 0 2 2]
//   proto info: {}
func (t *GeneratedCodeInfo_Annotation) GetBegin() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetBegin()

}

func (t *GeneratedCodeInfo_Annotation) Begin() int32 {
	return t.GetBegin()
}

// Identifies the ending offset in bytes in the generated code that
// relates to the identified offset. The end offset should be one past
// the last relevant byte (so the length of the text = end - begin).
//   See descriptorpb.GeneratedCodeInfo_Annotation GeneratedCodeInfo_Annotation.End
//   SourceCodeInfo.Location.Path: [4 20 3 0 2 3]
//   proto info: {}
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

/*
func (t *GeneratedCodeInfo_Annotation) MarshalJSON() (b []byte,err error) {
    if t.Empty() {
        return
    }
    buf := bytes.NewBuffer(nil)
    err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
    return buf.Bytes(), err
}
*/

// ExportFields returns can export fields
func (t *GeneratedCodeInfo_Annotation) ExportFields() map[string]interface{} {
	return map[string]interface{}{
		"Path":       t.Path(),
		"SourceFile": t.SourceFile(),
		"Begin":      t.Begin(),
		"End":        t.End(),
	}
}

// implement DescriptorCommon.Empty()
func (t *GeneratedCodeInfo_Annotation) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *GeneratedCodeInfo_Annotation) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *GeneratedCodeInfo_Annotation) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *GeneratedCodeInfo_Annotation) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
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
