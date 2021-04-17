// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Range of reserved numeric values. Reserved values may not be used by
// entries in the same enum. Reserved ranges may not overlap.
//
// Note that this is distinct from DescriptorProto.ReservedRange in that it
// is inclusive such that it can appropriately represent the entire int32
// domain.
//   See: descriptorpb.EnumDescriptorProto_EnumReservedRange
type EnumDescriptorProto_EnumReservedRange struct {
	common

	pb *descriptorpb.EnumDescriptorProto_EnumReservedRange
}

func newEnumDescriptorProto_EnumReservedRange(desc *descriptorpb.EnumDescriptorProto_EnumReservedRange) *EnumDescriptorProto_EnumReservedRange {
	t := new(EnumDescriptorProto_EnumReservedRange)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

//
// Inclusive.
//   See descriptorpb.EnumDescriptorProto_EnumReservedRange EnumDescriptorProto_EnumReservedRange.Start
//   SourceCodeInfo.Location.Path: [4 6 3 0 2 0]
//   proto info: {}
func (t *EnumDescriptorProto_EnumReservedRange) GetStart() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetStart()

}

func (t *EnumDescriptorProto_EnumReservedRange) Start() int32 {
	return t.GetStart()
}

//
// Inclusive.
//   See descriptorpb.EnumDescriptorProto_EnumReservedRange EnumDescriptorProto_EnumReservedRange.End
//   SourceCodeInfo.Location.Path: [4 6 3 0 2 1]
//   proto info: {}
func (t *EnumDescriptorProto_EnumReservedRange) GetEnd() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetEnd()

}

func (t *EnumDescriptorProto_EnumReservedRange) End() int32 {
	return t.GetEnd()
}

func (t *EnumDescriptorProto_EnumReservedRange) PbDescriptor() *descriptorpb.EnumDescriptorProto_EnumReservedRange {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *EnumDescriptorProto_EnumReservedRange) EnumDescriptorProto_EnumReservedRange() *descriptorpb.EnumDescriptorProto_EnumReservedRange {
	return t.PbDescriptor()
}

/*
func (t *EnumDescriptorProto_EnumReservedRange) MarshalJSON() (b []byte,err error) {
    if t.Empty() {
        return
    }
    buf := bytes.NewBuffer(nil)
    err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
    return buf.Bytes(), err
}
*/

// ExportFields returns can export fields
func (t *EnumDescriptorProto_EnumReservedRange) ExportFields() map[string]interface{} {
	return map[string]interface{}{
		"Start": t.Start(),
		"End":   t.End(),
	}
}

// implement DescriptorCommon.Empty()
func (t *EnumDescriptorProto_EnumReservedRange) Empty() bool {
	return t == nil || t.pb == nil
}

// implement DescriptorCommon.Index()
func (t *EnumDescriptorProto_EnumReservedRange) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

// implement DescriptorCommon.File()
func (t *EnumDescriptorProto_EnumReservedRange) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

// implement DescriptorCommon.Parent()
func (t *EnumDescriptorProto_EnumReservedRange) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

// implement DescriptorCommon.LocationPath()
func (t *EnumDescriptorProto_EnumReservedRange) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *EnumDescriptorProto_EnumReservedRange) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
