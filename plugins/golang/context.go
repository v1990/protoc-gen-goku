package golang

import (
	"fmt"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"strings"
)

var (
	globalPackages = map[string]*GoPackage{}
)

type Context struct {
	ctx *goku.Context
}

func NewContext(ctx *goku.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}

func (c Context) FuncMap() goku.FuncMap {
	return goku.FuncMap{
		"GoType":   c.GoType,
		"GoImport": c.GoImport,
	}
}

func (c Context) GoType(value interface{}) GoType {
	t := GoType{}

	switch v := value.(type) {
	case *descriptors.DescriptorProto:
		t.Base = helper.CamelCase(strings.Join(c.TypeNames(v), "_"))
	case *descriptors.EnumDescriptorProto:
		t.Base = helper.CamelCase(strings.Join(c.TypeNames(v), "_"))
	case *descriptors.EnumValueDescriptorProto:
		// TODO
	case *descriptors.FieldDescriptorProto:
		t.repeated = v.IsRepeated()
		if s, ok := baseTypes[v.GetType()]; ok {
			t.Base = s
		} else {
			// FIXME 更加通用
			name := v.GetTypeName()
			name = strings.TrimPrefix(name, ".google.protobuf.")
			name = strings.ReplaceAll(name, ".", "_")
			t.Base = name
			switch v.Type() {
			case descriptors.FieldDescriptorProto_TYPE_MESSAGE:
				t.pointer = true
			}
		}
	case string: // e.g. method.inputType/outputType
	}
	return t
}

func (c Context) TypeNames(d descriptors.Nestable) []string {
	var names []string
	names = append(names, d.GetName())
	for parent := d.ParentMessage(); parent != nil; parent = parent.ParentMessage() {
		name := parent.GetName()
		if len(name) == 0 {
			break
		}
		names = append(names, name)
	}
	names = helper.ReverseStrings(names)
	return names
}

func (c Context) GoImport(pkg string) *GoPackage {
	if p := c.getRegisteredPackageName(pkg); p != nil {
		return p
	}

	name := baseName(pkg)
	if IsGoKeyword(name) || IsGoPredeclaredIdentifier(name) {
		name += "_"
	}

	{
		i := 0
		nameBase := name
		for {
			if _, ok := globalPackages[name]; !ok {
				break
			}
			i++
			name = fmt.Sprintf("%s_%d", nameBase, i)
		}
	}

	p := &GoPackage{
		Name:       name,
		ImportPath: pkg,
	}

	globalPackages[p.Name] = p
	return p
}

func (c Context) getRegisteredPackageName(pkg string) *GoPackage {
	for _, p := range globalPackages {
		if p.ImportPath == pkg {
			return p
		}
	}
	return nil
}

// baseName returns the last path element of the name, with the last dotted suffix removed.
func baseName(name string) string {
	// First, find the last element
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	// Now drop the suffix
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[0:i]
	}
	return name
}
