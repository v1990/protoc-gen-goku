package descriptors

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FileOptions struct {
	common

	pb *descriptorpb.FileOptions

	uninterpreted_option []*UninterpretedOption
}

func newFileOptions(desc *descriptorpb.FileOptions) *FileOptions {
	t := new(FileOptions)
	t.pb = desc

	t.setDescriptor(t)

	return t
}

func (t *FileOptions) GetJavaPackage() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaPackage()

}

func (t *FileOptions) JavaPackage() string {
	return t.GetJavaPackage()
}

func (t *FileOptions) GetJavaOuterClassname() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaOuterClassname()

}

func (t *FileOptions) JavaOuterClassname() string {
	return t.GetJavaOuterClassname()
}

func (t *FileOptions) GetJavaMultipleFiles() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaMultipleFiles()

}

func (t *FileOptions) JavaMultipleFiles() bool {
	return t.GetJavaMultipleFiles()
}

func (t *FileOptions) GetJavaGenerateEqualsAndHash() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaGenerateEqualsAndHash()

}

func (t *FileOptions) JavaGenerateEqualsAndHash() bool {
	return t.GetJavaGenerateEqualsAndHash()
}

func (t *FileOptions) GetJavaStringCheckUtf8() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaStringCheckUtf8()

}

func (t *FileOptions) JavaStringCheckUtf8() bool {
	return t.GetJavaStringCheckUtf8()
}

func (t *FileOptions) GetOptimizeFor() (ret FileOptions_OptimizeMode) {
	if t.Empty() {
		return
	}

	return t.pb.GetOptimizeFor()

}

func (t *FileOptions) OptimizeFor() FileOptions_OptimizeMode {
	return t.GetOptimizeFor()
}

func (t *FileOptions) GetGoPackage() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetGoPackage()

}

func (t *FileOptions) GoPackage() string {
	return t.GetGoPackage()
}

func (t *FileOptions) GetCcGenericServices() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetCcGenericServices()

}

func (t *FileOptions) CcGenericServices() bool {
	return t.GetCcGenericServices()
}

func (t *FileOptions) GetJavaGenericServices() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetJavaGenericServices()

}

func (t *FileOptions) JavaGenericServices() bool {
	return t.GetJavaGenericServices()
}

func (t *FileOptions) GetPyGenericServices() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetPyGenericServices()

}

func (t *FileOptions) PyGenericServices() bool {
	return t.GetPyGenericServices()
}

func (t *FileOptions) GetPhpGenericServices() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetPhpGenericServices()

}

func (t *FileOptions) PhpGenericServices() bool {
	return t.GetPhpGenericServices()
}

func (t *FileOptions) GetDeprecated() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetDeprecated()

}

func (t *FileOptions) Deprecated() bool {
	return t.GetDeprecated()
}

func (t *FileOptions) GetCcEnableArenas() (ret bool) {
	if t.Empty() {
		return
	}

	return t.pb.GetCcEnableArenas()

}

func (t *FileOptions) CcEnableArenas() bool {
	return t.GetCcEnableArenas()
}

func (t *FileOptions) GetObjcClassPrefix() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetObjcClassPrefix()

}

func (t *FileOptions) ObjcClassPrefix() string {
	return t.GetObjcClassPrefix()
}

func (t *FileOptions) GetCsharpNamespace() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetCsharpNamespace()

}

func (t *FileOptions) CsharpNamespace() string {
	return t.GetCsharpNamespace()
}

func (t *FileOptions) GetSwiftPrefix() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetSwiftPrefix()

}

func (t *FileOptions) SwiftPrefix() string {
	return t.GetSwiftPrefix()
}

func (t *FileOptions) GetPhpClassPrefix() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetPhpClassPrefix()

}

func (t *FileOptions) PhpClassPrefix() string {
	return t.GetPhpClassPrefix()
}

func (t *FileOptions) GetPhpNamespace() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetPhpNamespace()

}

func (t *FileOptions) PhpNamespace() string {
	return t.GetPhpNamespace()
}

func (t *FileOptions) GetPhpMetadataNamespace() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetPhpMetadataNamespace()

}

func (t *FileOptions) PhpMetadataNamespace() string {
	return t.GetPhpMetadataNamespace()
}

func (t *FileOptions) GetRubyPackage() (ret string) {
	if t.Empty() {
		return
	}

	return t.pb.GetRubyPackage()

}

func (t *FileOptions) RubyPackage() string {
	return t.GetRubyPackage()
}

func (t *FileOptions) GetUninterpretedOption() (ret []*UninterpretedOption) {
	if t.Empty() {
		return
	}

	if t.uninterpreted_option != nil {
		return t.uninterpreted_option
	}

	t.uninterpreted_option = make([]*UninterpretedOption, len(t.pb.GetUninterpretedOption()))

	for i, item := range t.pb.GetUninterpretedOption() {
		elem := newUninterpretedOption(item)
		elem.setParent(t)
		elem.setIndex(i)
		t.uninterpreted_option[i] = elem
	}

	return t.uninterpreted_option

}

func (t *FileOptions) UninterpretedOption() []*UninterpretedOption {
	return t.GetUninterpretedOption()
}

func (t *FileOptions) PbDescriptor() *descriptorpb.FileOptions {
	if t == nil || t.pb == nil {
		return nil
	}
	return t.pb
}

func (t *FileOptions) FileOptions() *descriptorpb.FileOptions {
	return t.PbDescriptor()
}

func (t *FileOptions) MarshalJSON() (b []byte, err error) {
	if t.Empty() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = (&jsonpb.Marshaler{}).Marshal(buf, t.pb)
	return buf.Bytes(), err
}

func (t *FileOptions) Empty() bool {
	return t == nil || t.pb == nil
}

func (t *FileOptions) Index() int {
	if t.Empty() {
		return -1
	}

	return t.getIndex()
}

func (t *FileOptions) File() *FileDescriptorProto {
	if t.Empty() {
		return nil
	}

	return t.getFile()
}

func (t *FileOptions) Parent() DescriptorCommon {
	if t.Empty() {
		return nil
	}

	return t.getParent()
}

func (t *FileOptions) LocationPath() LocationPath {
	if t.Empty() {
		return nil
	}

	return t.getLocationPath()
}

func (t *FileOptions) Comments() *SourceCodeInfo_Location {
	if t.Empty() {
		return nil
	}
	return t.getComments()
}
