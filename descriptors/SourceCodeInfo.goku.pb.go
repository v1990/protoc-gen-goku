// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	bytes "bytes"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// ===================================================================
// Optional source code info  // Encapsulates information about the original source file from which a
// FileDescriptorProto was generated.
//   See: descriptorpb.SourceCodeInfo
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

// A Location identifies a piece of source code in a .proto file which
// corresponds to a particular definition.  This information is intended
// to be useful to IDEs, code indexers, documentation generators, and similar
// tools.
//
// For example, say we have a file like:
// message Foo {
// optional string foo = 1;
// }
// Let's look at just the field definition:
// optional string foo = 1;
// ^       ^^     ^^  ^  ^^^
// a       bc     de  f  ghi
// We have the following locations:
// span   path               represents
// [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
// [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
// [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
// [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
// [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).
//
// Notes:
// - A location may refer to a repeated field itself (i.e. not to any
// particular index within it).  This is used whenever a set of elements are
// logically enclosed in a single code segment.  For example, an entire
// extend block (possibly containing multiple extension definitions) will
// have an outer location whose path refers to the "extensions" repeated
// field without an index.
// - Multiple locations may have the same path.  This happens when a single
// logical declaration is spread out across multiple places.  The most
// obvious example is the "extend" block again -- there may be multiple
// extend blocks in the same scope, each of which will have the same path.
// - A location's span is not always a subset of its parent's span.  For
// example, the "extendee" of an extension declaration appears at the
// beginning of the "extend" block and is shared by all extensions within
// the block.
// - Just because a location's span is a subset of some other location's span
// does not mean that it is a descendant.  For example, a "group" defines
// both a type and a field in a single declaration.  Thus, the locations
// corresponding to the type and field and their components will overlap.
// - Code which tries to interpret locations should probably be designed to
// ignore those that it doesn't understand, as more types of locations could
// be recorded in the future.
//   See descriptorpb.SourceCodeInfo SourceCodeInfo.Location
//   SourceCodeInfo.Location.Path: [4 19 2 0]
//   proto info: {"name":"location","number":1,"label":"LABEL_REPEATED","type":"TYPE_MESSAGE","typeName":".google.protobuf.SourceCodeInfo.Location","jsonName":"location"}
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

// implement DescriptorCommon.Empty()
func (t *SourceCodeInfo) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *SourceCodeInfo) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *SourceCodeInfo) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *SourceCodeInfo) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *SourceCodeInfo) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}
