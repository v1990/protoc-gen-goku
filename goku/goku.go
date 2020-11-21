// protoc-gen-goku Generator
package goku

import (
	"flag"
	"github.com/Masterminds/sprig"
	"github.com/v1990/protoc-gen-goku/helper"
	"io"
	"log"
)

var (
	// 全局注入变量
	globalData = make(Data)
	// 全局注入函数
	globalFuncMap = make(FuncMap)

	defConf *Config
)

func init() {
	// 第三方函数 github.com/Masterminds/sprig
	RegisterFuncMap(FuncMap(sprig.TxtFuncMap()))
	// 基本函数
	RegisterFuncMap(FuncMap(helper.Functions))
}

// RegisterData 注册全局变量
func RegisterData(data Data) {
	for k, v := range data {
		globalData[k] = v
	}
}

// RegisterFuncMap 注入全局函数
func RegisterFuncMap(funcMap FuncMap) {
	for k, v := range funcMap {
		RegisterFunc(k, v)
	}
}

// RegisterFunc 注入全局函数
func RegisterFunc(name string, function interface{}) {
	globalFuncMap[name] = function
}

func Run(STDIN io.Reader, STDOUT io.Writer, STDERR io.Writer) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// 日志需要重定向的stderr
	log.SetOutput(STDERR)
	log.SetPrefix("[protoc-gen-goku]: ")
	log.SetFlags(log.Lshortfile)
	//log.SetFlags(log.Llongfile)

	g := NewGenerator()
	g.Run(STDIN, STDOUT, STDERR)

}
