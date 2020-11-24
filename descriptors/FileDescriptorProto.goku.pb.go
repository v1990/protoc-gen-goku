package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FileDescriptorProto struct {
	common

	pb *descriptorpb.FileDescriptorProto

	message_type []*DescriptorProto

	enum_type []*EnumDescriptorProto

	service []*ServiceDescriptorProto

	extension []*FieldDescriptorProto

	options *FileOptions

	source_code_info *SourceCodeInfo

	locByPath map[string]*SourceCodeInfo_Location
}

func newFileDescriptorProto(desc *descriptorpb.FileDescriptorProto) *FileDescriptorProto {
	t := new(FileDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *FileDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *FileDescriptorProto) Name() string {
	return t.GetName()
}

func (t *FileDescriptorProto) GetPackage() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetPackage()

}

func (t *FileDescriptorProto) Package() string {
	return t.GetPackage()
}

func (t *FileDescriptorProto) GetDependency() (ret []string) {
	if t.Empty() {
		return
	}

	return t.pb.GetDependency()

}

func (t *FileDescriptorProto) Dependency() []string {
	return t.GetDependency()
}

func (t *FileDescriptorProto) GetPublicDependency() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetPublicDependency()

}

func (t *FileDescriptorProto) PublicDependency() []int32 {
	return t.GetPublicDependency()
}

func (t *FileDescriptorProto) GetWeakDependency() (ret []int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetWeakDependency()

}

func (t *FileDescriptorProto) WeakDependency() []int32 {
	return t.GetWeakDependency()
}

func (t *FileDescriptorProto) GetMessageType() (ret []*DescriptorProto) {
	if t.Empty() {
		return
	}

	if t.message_type != nil {
		return t.message_type
	}

	t.message_type = make([]*DescriptorProto, len(t.pb.GetMessageType()))

	for i, item := range t.pb.GetMessageType() {
		elem := newDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.message_type[i] = elem
	}

	return t.message_type

}

func (t *FileDescriptorProto) MessageType() []*DescriptorProto {
	return t.GetMessageType()
}

func (t *FileDescriptorProto) GetEnumType() (ret []*EnumDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.enum_type != nil {
		return t.enum_type
	}

	t.enum_type = make([]*EnumDescriptorProto, len(t.pb.GetEnumType()))

	for i, item := range t.pb.GetEnumType() {
		elem := newEnumDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.enum_type[i] = elem
	}

	return t.enum_type

}

func (t *FileDescriptorProto) EnumType() []*EnumDescriptorProto {
	return t.GetEnumType()
}

func (t *FileDescriptorProto) GetService() (ret []*ServiceDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.service != nil {
		return t.service
	}

	t.service = make([]*ServiceDescriptorProto, len(t.pb.GetService()))

	for i, item := range t.pb.GetService() {
		elem := newServiceDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.service[i] = elem
	}

	return t.service

}

func (t *FileDescriptorProto) Service() []*ServiceDescriptorProto {
	return t.GetService()
}

func (t *FileDescriptorProto) GetExtension() (ret []*FieldDescriptorProto) {
	if t.Empty() {
		return
	}

	if t.extension != nil {
		return t.extension
	}

	t.extension = make([]*FieldDescriptorProto, len(t.pb.GetExtension()))

	for i, item := range t.pb.GetExtension() {
		elem := newFieldDescriptorProto(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.extension[i] = elem
	}

	return t.extension

}

func (t *FileDescriptorProto) Extension() []*FieldDescriptorProto {
	return t.GetExtension()
}

func (t *FileDescriptorProto) GetOptions() (ret *FileOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newFileOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *FileDescriptorProto) Options() *FileOptions {
	return t.GetOptions()
}

func (t *FileDescriptorProto) GetSourceCodeInfo() (ret *SourceCodeInfo) {
	if t.Empty() {
		return
	}

	if t.source_code_info != nil {
		return t.source_code_info
	}

	t.source_code_info = newSourceCodeInfo(t.pb.GetSourceCodeInfo())
	t.source_code_info.setParent(t)

	return t.source_code_info

}

func (t *FileDescriptorProto) SourceCodeInfo() *SourceCodeInfo {
	return t.GetSourceCodeInfo()
}

func (t *FileDescriptorProto) GetSyntax() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetSyntax()

}

func (t *FileDescriptorProto) Syntax() string {
	return t.GetSyntax()
}

func (t *FileDescriptorProto) PbDescriptor() *descriptorpb.FileDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *FileDescriptorProto) FileDescriptorProto() *descriptorpb.FileDescriptorProto {
	return t.PbDescriptor()
}

func (t *FileDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *FileDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *FileDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *FileDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *FileDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *FileDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *FileDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
