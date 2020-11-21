package goku

import (
	"fmt"
	"sort"
	"strings"
)

func (c *Context) baseFuncMap() FuncMap {
	return FuncMap{
		"Context": func() *Context {
			return c
		},
		"showFunc": c.ShowFunc,

		"getObject": c.GetObject,
	}
}

// 输出当前 Context 支持的函数
func (c *Context) ShowFunc() string {
	lines := make([]string, 0)
	for k, v := range c.FuncMap() {
		line := fmt.Sprintf("  %s: %T ", k, v)
		lines = append(lines, line)
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}
