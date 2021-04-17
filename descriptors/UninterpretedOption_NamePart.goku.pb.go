// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// The name of the uninterpreted option.  Each string represents a segment in
// a dot-separated name.  is_extension is true iff a segment represents an
// extension (denoted with parentheses in options specs in .proto files).
// E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
// "foo.(bar.baz).qux".
//   See: descriptorpb.UninterpretedOption_NamePart
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

//   See descriptorpb.UninterpretedOption_NamePart UninterpretedOption_NamePart.NamePart
//   SourceCodeInfo.Location.Path: [4 18 3 0 2 0]
//   proto info: {}
func (t *UninterpretedOption_NamePart) GetNamePart() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetNamePart()

}

func (t *UninterpretedOption_NamePart) NamePart() string {
	return t.GetNamePart()
}

//   See descriptorpb.UninterpretedOption_NamePart UninterpretedOption_NamePart.IsExtension
//   SourceCodeInfo.Location.Path: [4 18 3 0 2 1]
//   proto info: {}
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

/*
func (t *UninterpretedOption_NamePart) MarshalJSON() (b []byte,err error) {
    if t.Empty() {
        return
    }
    buf := bytes.NewBuffer(nil)
    err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
    return buf.Bytes(), err
}
*/

// ExportFields returns can export fields
func (t *UninterpretedOption_NamePart) ExportFields() map[string]interface{} {
	return map[string]interface{}{
		"NamePart":    t.NamePart(),
		"IsExtension": t.IsExtension(),
	}
}

// implement DescriptorCommon.Empty()
func (t *UninterpretedOption_NamePart) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *UninterpretedOption_NamePart) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *UninterpretedOption_NamePart) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *UninterpretedOption_NamePart) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
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
