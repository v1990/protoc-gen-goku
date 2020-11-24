package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FieldDescriptorProto struct {
	common

	pb *descriptorpb.FieldDescriptorProto

	options *FieldOptions
}

func newFieldDescriptorProto(desc *descriptorpb.FieldDescriptorProto) *FieldDescriptorProto {
	t := new(FieldDescriptorProto)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *FieldDescriptorProto) GetName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetName()

}

func (t *FieldDescriptorProto) Name() string {
	return t.GetName()
}

func (t *FieldDescriptorProto) GetNumber() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetNumber()

}

func (t *FieldDescriptorProto) Number() int32 {
	return t.GetNumber()
}

func (t *FieldDescriptorProto) GetLabel() (ret FieldDescriptorProto_Label) {
	if t.Empty() {
		return
	}

	return t.pb.GetLabel()

}

func (t *FieldDescriptorProto) Label() FieldDescriptorProto_Label {
	return t.GetLabel()
}

func (t *FieldDescriptorProto) GetType() (ret FieldDescriptorProto_Type) {
	if t.Empty() {
		return
	}

	return t.pb.GetType()

}

func (t *FieldDescriptorProto) Type() FieldDescriptorProto_Type {
	return t.GetType()
}

func (t *FieldDescriptorProto) GetTypeName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetTypeName()

}

func (t *FieldDescriptorProto) TypeName() string {
	return t.GetTypeName()
}

func (t *FieldDescriptorProto) GetExtendee() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetExtendee()

}

func (t *FieldDescriptorProto) Extendee() string {
	return t.GetExtendee()
}

func (t *FieldDescriptorProto) GetDefaultValue() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetDefaultValue()

}

func (t *FieldDescriptorProto) DefaultValue() string {
	return t.GetDefaultValue()
}

func (t *FieldDescriptorProto) GetOneofIndex() (ret int32) {
	if t.Empty() {
		return
	}

	return t.pb.GetOneofIndex()

}

func (t *FieldDescriptorProto) OneofIndex() int32 {
	return t.GetOneofIndex()
}

func (t *FieldDescriptorProto) GetJsonName() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetJsonName()

}

func (t *FieldDescriptorProto) JsonName() string {
	return t.GetJsonName()
}

func (t *FieldDescriptorProto) GetOptions() (ret *FieldOptions) {
	if t.Empty() {
		return
	}

	if t.options != nil {
		return t.options
	}

	t.options = newFieldOptions(t.pb.GetOptions())
	t.options.setParent(t)

	return t.options

}

func (t *FieldDescriptorProto) Options() *FieldOptions {
	return t.GetOptions()
}

func (t *FieldDescriptorProto) PbDescriptor() *descriptorpb.FieldDescriptorProto {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *FieldDescriptorProto) FieldDescriptorProto() *descriptorpb.FieldDescriptorProto {
	return t.PbDescriptor()
}

func (t *FieldDescriptorProto) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *FieldDescriptorProto) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *FieldDescriptorProto) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *FieldDescriptorProto) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *FieldDescriptorProto) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *FieldDescriptorProto) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *FieldDescriptorProto) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
