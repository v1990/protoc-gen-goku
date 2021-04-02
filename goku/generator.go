package goku

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"github.com/v1990/protoc-gen-goku/helper"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	pluginpb "google.golang.org/protobuf/types/pluginpb"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Generator struct {
	conf   Config
	params map[string]string

	ctx *Context

	request  *pluginpb.CodeGeneratorRequest  // The input.
	response *pluginpb.CodeGeneratorResponse // The output.

	// name: google/protobuf/empty.proto
	filesByName map[string]*descriptors.FileDescriptorProto

	// pb类型名称与实例的映射
	// key: 完整的类型名称 .google.protobuf.FieldDescriptorProto.Label
	// value: message(descriptors.DescriptorProto) or enum(descriptors.EnumDescriptorProto)
	typesByName map[string]descriptors.ProtoType

	debug bool
}

func NewGenerator() *Generator {
	g := new(Generator)
	g.params = make(map[string]string)
	g.filesByName = make(map[string]*descriptors.FileDescriptorProto)
	g.typesByName = make(map[string]descriptors.ProtoType)
	g.ctx = newContext(g)
	g.request = new(pluginpb.CodeGeneratorRequest)
	g.response = new(pluginpb.CodeGeneratorResponse)
	return g
}

func (g *Generator) Run(IN io.Reader, OUT io.Writer, ERR io.Writer) {

	data, err := ioutil.ReadAll(IN)
	g.FatalOnErr(err, "reading input")

	ioutil.WriteFile("request.pb", data, 0644)

	err = proto.Unmarshal(data, g.request)
	g.FatalOnErr(err, "parsing input proto")

	if len(g.request.FileToGenerate) == 0 {
		log.Fatalln("no files to generateFile")
	}

	g.parseParameters(g.request.GetParameter())
	g.conf = g.getConfig()

	g.logRequest(g.request)

	// ==================================
	g.wrapAllFiles(g.request.GetProtoFile())
	g.initPlugins()
	g.generateFiles(g.request.GetFileToGenerate())
	// ==================================

	data, err = proto.Marshal(g.response)
	g.FatalOnErr(err, "failed to marshal output proto")
	_, err = OUT.Write(data)
	g.FatalOnErr(err, "failed to write output proto")

}

func (g *Generator) parseParameters(parameter string) {
	for _, p := range strings.Split(parameter, ",") {
		if i := strings.Index(p, "="); i < 0 {
			g.params[p] = ""
		} else {
			g.params[p[0:i]] = p[i+1:]
		}
	}

	for k, v := range g.params {
		switch k {
		case "debug":
			g.debug, _ = strconv.ParseBool(v)
		}
	}

	g.setDefaultParam("workPath", ".")
	g.setDefaultParam("outPath", "")
	g.setDefaultParam("templatePath", "templates")
	g.setDefaultParam("conf", "goku.yaml")

}

func (g *Generator) wrapAllFiles(files []*descriptorpb.FileDescriptorProto) {
	for _, f := range files {
		if _, ok := g.filesByName[f.GetName()]; ok {
			continue
		}
		file := descriptors.NewFileDescriptorProto(f)

		g.filesByName[file.GetName()] = file

		for _, m := range file.GetMessageType() {
			g.recordType(m)
			for _, mm := range m.GetNestedType() {
				g.recordType(mm)
			}
			for _, mm := range m.GetEnumType() {
				g.recordType(mm)
			}
		}
		for _, m := range file.GetEnumType() {
			g.recordType(m)
		}
	}
}
func (g *Generator) recordType(t descriptors.ProtoType) {
	name := t.ProtoType().FullTypeName()
	g.typesByName[name] = t
	//g.Debug("recordType: %-60s ==> [%T]%s", name, t, t.GetName())
}

func (g *Generator) GetObject(typeName string) descriptors.ProtoType {
	return g.MustGetDescriptorByName(typeName)
}

func (g *Generator) GetFileByName(protoFileName string) *descriptors.FileDescriptorProto {
	return g.filesByName[protoFileName]
}

func (g *Generator) GetGlobalCtx() *Context {
	return g.ctx
}

// 通过protoc格式的名称 获取 proto 对象；
//  @param typeName .google.protobuf.FileDescriptorProto
//  @return
//  - message: *descriptors.DescriptorProto
//  - enum: *descriptors.EnumDescriptorProto
func (g *Generator) GetDescriptorByName(typeName string) descriptors.ProtoType {
	return g.typesByName[typeName]
}
func (g *Generator) MustGetDescriptorByName(typeName string) descriptors.ProtoType {
	if p := g.GetDescriptorByName(typeName); p != nil {
		return p
	}
	g.Fatal("can not found type: %s", typeName)
	return nil
}
func (g *Generator) EachType(f func(name string, obj descriptors.ProtoType)) {
	for name, obj := range g.typesByName {
		f(name, obj)
	}
}

func (g *Generator) initPlugins() {
	for _, p := range plugins {
		p.Init(g)
	}
}

func (g *Generator) generateFiles(filenames []string) {
	gCtx := g.ctx.WithLoop(LoopOnce, nil)
	g.executeJobs(gCtx)

	for _, filename := range filenames {
		file := g.filesByName[filename]
		g.generateFile(file)
	}
}

// 输出文件
// See pluginpb.CodeGeneratorResponse_File
func (g *Generator) WriteOutFile(filename string, content []byte) error {
	// 默认输出到 stderr
	switch filename {
	case "", "stderr", "stdin":
		_, err := fmt.Fprint(os.Stderr, string(content)+"\n")
		return err
	}

	// 如果是 绝对路径: 直接写文件
	if filepath.IsAbs(filename) {
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return err
		}
		if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
			return err
		}
		return nil
	}

	if g.params["outPath"] != "" && !strings.HasPrefix(filename, g.params["outPath"]) {
		filename = filepath.Join(g.params["outPath"], filename)
		return g.WriteOutFile(filename, content)
	}

	// 写到 protoc 的 response
	sContent := string(content)
	outFile := &pluginpb.CodeGeneratorResponse_File{
		Name:           &filename,
		InsertionPoint: nil,
		Content:        &sContent,
	}
	g.response.File = append(g.response.File, outFile)

	return nil
}

func (g *Generator) populateCtx(ctx *Context) {
	ctx.populate()
}

// See Context.Value
func (g *Generator) Value(key interface{}) interface{} {
	switch k := key.(type) {
	case string:
		for _, data := range g.conf.Data {
			if v, ok := data[k]; ok {
				return v
			}
		}
		// FIXME:
		//if v, ok := g.conf.Data[k]; ok {
		//	return v
		//}
		if v, ok := g.params[k]; ok {
			return v
		}
	}
	return nil
}
func (g *Generator) Param(key string, def string) string {
	if v, ok := g.params[key]; ok {
		return v
	}
	return def
}
func (g *Generator) setDefaultParam(key string, def string) string {
	g.params[key] = g.Param(key, def)
	return g.params[key]
}

func (g *Generator) logRequest(r *pluginpb.CodeGeneratorRequest) {
	jsonStr := helper.ShowJSON(r, 3)
	//jsonStr, err := (&jsonpb.Marshaler{
	//	Indent: " ",
	//}).MarshalToString(r)
	//if err != nil {
	//	return
	//}
	logFile := "request.json"
	ioutil.WriteFile(logFile, []byte(jsonStr), 0644)
}
