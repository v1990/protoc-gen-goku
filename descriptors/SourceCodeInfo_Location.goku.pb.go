// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

//   See: descriptorpb.SourceCodeInfo_Location
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

// Identifies which part of the FileDescriptorProto was defined at this
// location.
//
// Each element is a field number or an index.  They form a path from
// the root FileDescriptorProto to the place where the definition.  For
// example, this path:
// [ 4, 3, 2, 7, 1 ]
// refers to:
// file.message_type(3)  // 4, 3
// .field(7)         // 2, 7
// .name()           // 1
// This is because FileDescriptorProto.message_type has field number 4:
// repeated DescriptorProto message_type = 4;
// and DescriptorProto.field has field number 2:
// repeated FieldDescriptorProto field = 2;
// and FieldDescriptorProto.name has field number 1:
// optional string name = 1;
//
// Thus, the above path gives the location of a field name.  If we removed
// the last element:
// [ 4, 3, 2, 7 ]
// this path refers to the whole field declaration (from the beginning
// of the label to the terminating semicolon).
//   See descriptorpb.SourceCodeInfo_Location SourceCodeInfo_Location.Path
//   SourceCodeInfo.Location.Path: [4 19 3 0 2 0]
//   proto info: {"name":"path","number":1,"label":"LABEL_REPEATED","type":"TYPE_INT32","jsonName":"path","options":{"packed":true}}
func (t *SourceCodeInfo_Location) GetPath() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetPath()

}

func (t *SourceCodeInfo_Location) Path() []int32 {
	return t.GetPath()
}

// Always has exactly three or four elements: start line, start column,
// end line (optional, otherwise assumed same as start line), end column.
// These are packed into a single field for efficiency.  Note that line
// and column numbers are zero-based -- typically you will want to add
// 1 to each before displaying to a user.
//   See descriptorpb.SourceCodeInfo_Location SourceCodeInfo_Location.Span
//   SourceCodeInfo.Location.Path: [4 19 3 0 2 1]
//   proto info: {"name":"span","number":2,"label":"LABEL_REPEATED","type":"TYPE_INT32","jsonName":"span","options":{"packed":true}}
func (t *SourceCodeInfo_Location) GetSpan() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetSpan()

}

func (t *SourceCodeInfo_Location) Span() []int32 {
	return t.GetSpan()
}

// If this SourceCodeInfo represents a complete declaration, these are any
// comments appearing before and after the declaration which appear to be
// attached to the declaration.
//
// A series of line comments appearing on consecutive lines, with no other
// tokens appearing on those lines, will be treated as a single comment.
//
// leading_detached_comments will keep paragraphs of comments that appear
// before (but not connected to) the current element. Each paragraph,
// separated by empty lines, will be one comment element in the repeated
// field.
//
// Only the comment content is provided; comment markers (e.g. //) are
// stripped out.  For block comments, leading whitespace and an asterisk
// will be stripped from the beginning of each line other than the first.
// Newlines are included in the output.
//
// Examples:
//
// optional int32 foo = 1;  // Comment attached to foo.
// Comment attached to bar.
// optional int32 bar = 2;
//
// optional string baz = 3;
// Comment attached to baz.
// Another line attached to baz.
//
// Comment attached to qux.
//
// Another line attached to qux.
// optional double qux = 4;
//
// Detached comment for corge. This is not leading or trailing comments
// to qux or corge because there are blank lines separating it from
// both.
//
// Detached comment for corge paragraph 2.
//
// optional string corge = 5;
//  Block comment attached
//  to corge.  Leading asterisks
//  will be removed. */
//  Block comment attached to
//  grault. */
// optional int32 grault = 6;
//
// ignored detached comments.
//   See descriptorpb.SourceCodeInfo_Location SourceCodeInfo_Location.LeadingComments
//   SourceCodeInfo.Location.Path: [4 19 3 0 2 2]
//   proto info: {"name":"leading_comments","number":3,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"leadingComments"}
func (t *SourceCodeInfo_Location) GetLeadingComments() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetLeadingComments()

}

func (t *SourceCodeInfo_Location) LeadingComments() string {
	return t.GetLeadingComments()
}

//   See descriptorpb.SourceCodeInfo_Location SourceCodeInfo_Location.TrailingComments
//   SourceCodeInfo.Location.Path: [4 19 3 0 2 3]
//   proto info: {"name":"trailing_comments","number":4,"label":"LABEL_OPTIONAL","type":"TYPE_STRING","jsonName":"trailingComments"}
func (t *SourceCodeInfo_Location) GetTrailingComments() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetTrailingComments()

}

func (t *SourceCodeInfo_Location) TrailingComments() string {
	return t.GetTrailingComments()
}

//   See descriptorpb.SourceCodeInfo_Location SourceCodeInfo_Location.LeadingDetachedComments
//   SourceCodeInfo.Location.Path: [4 19 3 0 2 4]
//   proto info: {"name":"leading_detached_comments","number":6,"label":"LABEL_REPEATED","type":"TYPE_STRING","jsonName":"leadingDetachedComments"}
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

// implement DescriptorCommon.Empty()
func (t *SourceCodeInfo_Location) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *SourceCodeInfo_Location) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *SourceCodeInfo_Location) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *SourceCodeInfo_Location) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *SourceCodeInfo_Location) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}
