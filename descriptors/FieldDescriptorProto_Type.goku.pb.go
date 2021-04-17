// Code generated by protoc-gen-goku. DO NOT EDIT.
// source: google/protobuf/descriptor.proto

package descriptors

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// FieldDescriptorProto_Type
//  See descriptorpb.FieldDescriptorProto_Type
type FieldDescriptorProto_Type = descriptorpb.FieldDescriptorProto_Type

const (
	// 0 is reserved for errors.
	// Order is weird for historical reasons.
	FieldDescriptorProto_TYPE_DOUBLE = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	FieldDescriptorProto_TYPE_FLOAT  = descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
	// negative values are likely.
	FieldDescriptorProto_TYPE_INT64  = descriptorpb.FieldDescriptorProto_TYPE_INT64
	FieldDescriptorProto_TYPE_UINT64 = descriptorpb.FieldDescriptorProto_TYPE_UINT64
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
	// negative values are likely.
	FieldDescriptorProto_TYPE_INT32   = descriptorpb.FieldDescriptorProto_TYPE_INT32
	FieldDescriptorProto_TYPE_FIXED64 = descriptorpb.FieldDescriptorProto_TYPE_FIXED64
	FieldDescriptorProto_TYPE_FIXED32 = descriptorpb.FieldDescriptorProto_TYPE_FIXED32
	FieldDescriptorProto_TYPE_BOOL    = descriptorpb.FieldDescriptorProto_TYPE_BOOL
	FieldDescriptorProto_TYPE_STRING  = descriptorpb.FieldDescriptorProto_TYPE_STRING
	// Tag-delimited aggregate.
	// Group type is deprecated and not supported in proto3. However, Proto3
	// implementations should still be able to parse the group wire format and
	// treat group fields as unknown fields.
	FieldDescriptorProto_TYPE_GROUP   = descriptorpb.FieldDescriptorProto_TYPE_GROUP
	FieldDescriptorProto_TYPE_MESSAGE = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE // Length-delimited aggregate.
	// New in version 2.
	FieldDescriptorProto_TYPE_BYTES    = descriptorpb.FieldDescriptorProto_TYPE_BYTES
	FieldDescriptorProto_TYPE_UINT32   = descriptorpb.FieldDescriptorProto_TYPE_UINT32
	FieldDescriptorProto_TYPE_ENUM     = descriptorpb.FieldDescriptorProto_TYPE_ENUM
	FieldDescriptorProto_TYPE_SFIXED32 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED32
	FieldDescriptorProto_TYPE_SFIXED64 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED64
	FieldDescriptorProto_TYPE_SINT32   = descriptorpb.FieldDescriptorProto_TYPE_SINT32 // Uses ZigZag encoding.
	FieldDescriptorProto_TYPE_SINT64   = descriptorpb.FieldDescriptorProto_TYPE_SINT64 // Uses ZigZag encoding.
)
