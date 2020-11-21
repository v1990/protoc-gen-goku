package goku

import (
	"github.com/ghodss/yaml"
	"github.com/v1990/protoc-gen-goku/helper"
	"io/ioutil"
	"strconv"
)

type Loop string

const (
	LoopFile          Loop = "file"
	LoopService       Loop = "service"
	LoopMethod        Loop = "method"
	LoopMessage       Loop = "message"
	LoopEnum          Loop = "enum"
	LoopNestedMessage Loop = "nested_message" // 嵌套在message中的message
	LoopNestedEnum    Loop = "nested_enum"    // 嵌套在message中的enum
)

type Config struct {
	// 全局变量 - 支持模版解析
	Data Data
	// 启用插件列表
	Plugins []string
	// 任务列表
	Jobs []Job
}

type Job struct {
	// 任务名称
	Name string
	// 启用：所有条件符合则启用
	Enable Enable
	// 根据所处阶段
	Loop LoopCondition
	// 用户自定义的判断条件：返回 "true"/"false"
	//  - 支持模版解析
	If IfCondition
	// 模版内容
	Template string
	// 模板路径 - 支持模板解析
	//  -- Template 为空时才会读取模板文件
	TemplatePath string
	// 输出文件路径
	Out string
	// 启用插件列表
	//
	Plugins []string
	// 任务级别的变量 - 支持模版解析
	Data Data
}

type Enable struct {
	// 根据所处阶段
	Loop LoopCondition
	// 用户自定义的判断条件：返回 "true"/"false"
	//  - 支持模版解析
	If IfCondition
}

type Condition interface {
	OK(ctx *Context) bool
}
type LoopCondition []string
type IfCondition string

func (t LoopCondition) OK(ctx *Context) bool {
	if len(t) == 0 {
		return true
	}
	return helper.InStrings(string(ctx.Loop()), t...)
}

func (t IfCondition) OK(ctx *Context) bool {
	if len(t) == 0 {
		return true
	}
	str := ctx.MustEval(string(t))
	ok, err := strconv.ParseBool(str)
	ctx.FatalOnErr(err, "parse if: %s parsed:%s", t, str)
	return ok
}

func (j Job) GetConditions() []Condition {
	return []Condition{
		j.Enable.Loop,
		j.Enable.If,
		j.Loop,
		j.If,
	}
}

func (j Job) IsEnable(ctx *Context) bool {
	for _, condition := range j.GetConditions() {
		if !condition.OK(ctx) {
			return false
		}
	}
	return true
}

//type IConfig interface {
//	GetVar() []Data
//	GetPlugin() []string
//}
//
//func (t Config) GetVar() []Data     { return t.Data }
//func (t Config) GetPlugin() []string { return t.Plugins }
//
//func (t Job) GetVar() []Data     { return t.Data }
//func (t Job) GetPlugin() []string { return t.Plugins }

//
//func (t Job) IsEnable(ctx *Context) bool {
//	return t.Enable.OK(ctx)
//}

//func (t Job) getOutPlugins() []string {
//	if len(t.OutPlugins) == 0 {
//		return t.Plugins
//	}
//
//	exclude := make([]string, 0)
//	include := make([]string, 0)
//
//	for _, name := range t.OutPlugins {
//		if strings.HasPrefix(name, "-") {
//			exclude = append(exclude, name[1:])
//		} else {
//			include = append(include, name)
//		}
//	}
//
//	for _, name := range t.Plugins {
//		// 已经被排除的，不要
//		if helper.InStrings(name, exclude...) {
//			continue
//		}
//		// 已经包含了，不要
//		if helper.InStrings(name, include...) {
//			continue
//		}
//
//		include = append(include, name)
//	}
//
//	return include
//}
//
//func (c Enable) OK(ctx *Context) bool {
//	return c.loopOK(ctx) && c.ifOK(ctx)
//}
//
//func (c Enable) loopOK(ctx *Context) bool {
//	if len(c.Loop) == 0 {
//		return true
//	}
//	//ctx.Debug("loopOK: %+v", utils.InStrings(string(ctx.Loop()), c.Loop...))
//	return helper.InStrings(string(ctx.Loop()), c.Loop...)
//}
//
//func (c Enable) ifOK(ctx *Context) bool {
//	if len(c.If) == 0 {
//		return true
//	}
//	str := ctx.parseTextTpl(c.If, ctx)
//	ok, err := strconv.ParseBool(str)
//	//log.Println("ifOK:",c.If,str,ok,err)
//	ctx.FatalOnErr(err, "parse if: %s parsed:%s", c.If, str)
//
//	return ok
//}

func (g *Generator) getConfig() Config {
	// TODO 先读取编译进来的配置对象

	// 读取配置文件
	filename := g.params["conf"]
	if len(filename) == 0 {
		filename = "config.yaml"
	}

	if !helper.FileExists(filename) {

	}

	body, err := ioutil.ReadFile(filename)
	g.FatalOnErr(err, "read config: %s", filename)

	var conf Config
	err = yaml.Unmarshal(body, &conf)
	g.FatalOnErr(err, "unmarshal conf[YAML]: %s", filename)

	// TODO 配置检查

	g.Debug("conf: %s", helper.ShowJSON(conf))
	return conf
}
