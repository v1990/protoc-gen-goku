package goku

import (
	"bytes"
	"fmt"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"github.com/v1990/protoc-gen-goku/helper"
	"strings"
	"text/template"
)

type Data map[string]interface{}

func (data Data) Copy() Data {
	out := make(Data)
	for k, v := range data {
		out[k] = v
	}
	return out
}
func (data Data) DoMerge(other Data) {
	for k, v := range other {
		data[k] = v
	}
}

type FuncMap template.FuncMap

func (m FuncMap) Copy() FuncMap {
	out := make(FuncMap)
	for k, v := range m {
		out[k] = v
	}
	return out
}

func (m FuncMap) DoMerge(other FuncMap) {
	for k, v := range other {
		m[k] = v
	}
}

type DescriptorObject interface {
	descriptors.DescriptorCommon
	GetName() string
}

type Context struct {
	*Generator
	globalData map[interface{}]interface{}

	data    Data
	funcMap FuncMap
	loop    Loop
	job     *Job

	content []byte

	object DescriptorObject

	file          *File
	service       *Service
	method        *Method
	message       *Message
	enumObj       *Enum
	parentMessage *Message
}

func newContext(g *Generator) *Context {
	return &Context{
		Generator:  g,
		globalData: make(map[interface{}]interface{}),
		data:       make(Data),
		funcMap:    make(FuncMap),
	}
}

func (c *Context) copy() *Context {
	nc := new(Context)
	*nc = *c

	// 这两个map要深拷贝
	nc.data = c.data.Copy()
	nc.funcMap = c.funcMap.Copy()

	return nc
}

func (c *Context) WithLoop(loop Loop, desc DescriptorObject) *Context {
	cc := c.copy()
	cc.loop = loop
	cc.object = desc

	switch d := desc.(type) {
	case *descriptors.FileDescriptorProto:
		cc.file = d
	case *descriptors.ServiceDescriptorProto:
		cc.service = d
	case *descriptors.DescriptorProto:
		cc.message = d
	case *descriptors.EnumDescriptorProto:
		cc.enumObj = d
	case *descriptors.MethodDescriptorProto:
		cc.method = d
	default:
	}

	if d, ok := desc.(descriptors.Nestable); ok {
		cc.parentMessage = d.ParentMessage()
	}

	return cc
}

func (c *Context) withJob(job Job) *Context {
	cc := c.copy()
	cc.job = &job
	return cc
}

// Loop returns current loop
func (c *Context) Loop() Loop {
	return c.loop
}

func (c *Context) MergeData(data Data) {
	c.data.DoMerge(data)
}

func (c *Context) MergeFuncMap(funcMap FuncMap) {
	c.funcMap.DoMerge(funcMap)
}

func (c *Context) Data() Data {
	return c.data
}

func (c *Context) Value(key interface{}) interface{} {
	switch k := key.(type) {
	case string:
		if v, ok := c.data[k]; ok {
			return v
		}
	}

	return c.globalData[key]
}
func (c *Context) SetGlobal(kvs ...interface{}) {
	for len(kvs) >= 2 {
		c.globalData[kvs[0]] = kvs[1]
		kvs = kvs[2:]
	}
}

func (c *Context) FuncMap() FuncMap {
	return c.funcMap
}
func (c *Context) tplFuncMap() template.FuncMap {
	return template.FuncMap(c.FuncMap())
}

func (c *Context) Object() DescriptorObject {
	return c.object
}

// File returns current FileDescriptorProto
func (c *Context) File() *File {
	return c.file
}

// Service returns current ServiceDescriptorProto. ONLY [ LoopService LoopMethod ]
func (c *Context) Service() *Service {
	return c.service
}

// Method returns current MethodDescriptorProto. ONLY [ LoopMethod ]
func (c *Context) Method() *Method {
	return c.method
}

// ParentMessage  ONLY [ LoopNestedMessage LoopNestedEnum ]
func (c *Context) ParentMessage() *Message {
	return c.Message().ParentMessage()
}

// Message ONLY [ LoopMessage ]
func (c *Context) Message() *Message {
	return c.message
}

// Enum ONLY [ LoopEnum ]
func (c *Context) Enum() *Enum {
	return c.enumObj
}

func (c *Context) SetContent(content []byte) {
	c.content = content
}

func (c *Context) Content() []byte {
	return c.content
}

func (c *Context) GetFileName() string {
	return c.file.Name()
}

func (c *Context) MustEval(text string, args ...interface{}) string {
	body, err := c.Eval(text, args...)
	c.FatalOnErr(err, "eval: %s", err)
	return body
}

// 执行一段模板（脚本）
func (c *Context) Eval(text string, args ...interface{}) (string, error) {
	data := c.Data()

	if len(args) > 0 {
		data = c.Data().Copy()

		for _, arg := range args {
			switch t := arg.(type) {
			case Data:
				data.DoMerge(t)
			default: // TODO
				//	data.DoMerge(helper.Struct2Map(arg))
			}
		}
	}

	tpl, err := template.New("text").Funcs(c.tplFuncMap()).Parse(text)
	if err != nil {
		return "", fmt.Errorf("eval.parse.err: %w  [text] %s", err, text)
	}

	buf := bytes.NewBuffer(nil)
	err = tpl.Execute(buf, data)
	if err != nil {
		return "", fmt.Errorf("eval:execute err: %w  [text] %s", err, text)
	}
	return buf.String(), nil
}

func (c *Context) callPlugins(f func(plugin Plugin)) {
	var include, exclude []string
	if c.job != nil {
		for _, name := range c.job.Plugins {
			if strings.HasPrefix(name, "-") {
				exclude = append(exclude, name[1:])
			} else {
				include = append(include, name)
			}
		}
	}

	include = append(include, c.conf.Plugins...)

	executed := make(map[string]bool)
	for _, pluginName := range include {
		if helper.InStrings(pluginName, exclude...) {
			continue
		}
		if executed[pluginName] {
			continue
		}
		executed[pluginName] = true
		if p, ok := plugins[pluginName]; ok {
			f(p)
		}
	}

}
func (c *Context) parseConfData() {
	parseConfData(c, c.conf.Data)
	if c.job != nil {
		parseConfData(c, c.job.Data)
	}
}

func parseConfData(ctx *Context, list []Data) {
	for _, data := range list {
		parseData(ctx, data)
	}
}
func parseData(ctx *Context, data Data) {
	for k, vvv := range data {
		switch v := vvv.(type) {
		case string:
			ctx.data[k] = ctx.MustEval(v)
		default:
			ctx.data[k] = v
		}
	}
}
