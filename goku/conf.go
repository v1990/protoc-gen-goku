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
	// 声明全局变量
	//  - 支持模版解析
	// 因为 map 是无序的，所以不可相互引用
	Data Data
	// 全局启用插件列表
	Plugins []string
	// 任务列表
	// TODO 改为 map[string]Job,并且params增加jobs，excludeJobs来选择要执行的job
	Jobs []Job
}

type Job struct {
	// 任务名称
	Name string
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
	// 任务级别的变量
	// - 与全局的 Config.Data
	Data Data
}

type Condition interface {
	OK(ctx *Context) bool
}
type LoopCondition []Loop
type IfCondition string

func (t LoopCondition) OK(ctx *Context) bool {
	if len(t) == 0 {
		return true
	}
	return helper.In(ctx.Loop(), t)
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
