package goku

import (
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"github.com/v1990/protoc-gen-goku/goku/internal"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Generator struct {
	// 以后要对这个 Generator 进行重构，所以将其作为内部包
	gen *internal.Generator

	conf   Config
	params map[string]string

	ctx *Context

	request  *plugin.CodeGeneratorRequest  // The input.
	response *plugin.CodeGeneratorResponse // The output.

	allFiles       []*descriptors.FileDescriptorProto
	allFilesByName map[string]*descriptors.FileDescriptorProto
}

func NewGenerator() *Generator {
	g := new(Generator)
	g.params = make(map[string]string)
	g.allFilesByName = make(map[string]*descriptors.FileDescriptorProto)
	g.ctx = newContext(g)
	g.request = new(plugin.CodeGeneratorRequest)
	g.response = new(plugin.CodeGeneratorResponse)
	return g
}

func (g *Generator) Run(IN io.Reader, OUT io.Writer, ERR io.Writer) {

	data, err := ioutil.ReadAll(IN)
	g.FatalOnErr(err, "reading input")

	err = proto.Unmarshal(data, g.request)
	g.FatalOnErr(err, "parsing input proto")

	if len(g.request.FileToGenerate) == 0 {
		log.Fatalln("no files to generateFile")
	}

	g.parseParameters(g.request.GetParameter())
	g.conf = g.getConfig()

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
}

func (g *Generator) wrapAllFiles(files []*descriptorpb.FileDescriptorProto) {
	for _, f := range files {
		file := descriptors.NewFileDescriptorProto(f)
		g.allFiles = append(g.allFiles, file)
		g.allFilesByName[file.GetName()] = file
	}
}

func (g *Generator) initPlugins() {
	for _, p := range plugins {
		p.Init(g)
	}
}

func (g *Generator) generateFiles(filenames []string) {
	for _, filename := range filenames {
		file := g.allFilesByName[filename]
		g.generateFile(file)
	}
}

func (g *Generator) addOutFile(filename string, content string) {
	g.response.File = append(g.response.File, &plugin.CodeGeneratorResponse_File{
		Name:           &filename,
		InsertionPoint: nil,
		Content:        &content,
	})
}

func (g *Generator) populateCtx(ctx *Context) {
	// 函数
	ctx.MergeFuncMap(globalFuncMap)
	ctx.MergeFuncMap(ctx.baseFuncMap())
	// 静态变量
	ctx.MergeData(globalData)
	ctx.MergeData(Data{
		"Params":        g.params,
		"Ctx":           ctx,
		"File":          ctx.File(),
		"Service":       ctx.Service(),
		"Message":       ctx.Message(),
		"Enum":          ctx.Enum(),
		"Object":        ctx.Object(),
		"ParentMessage": ctx.ParentMessage(),
	})
	// 启用插件
	ctx.callPlugins(func(plugin Plugin) {
		plugin.BeforeExecute(ctx)
	})
	// 解析配置数据
	ctx.parseConfData()

}
