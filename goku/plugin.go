package goku

// Plugin 模板数据插件
type Plugin interface {
	Name() string
	Init(c *Generator)
	// BeforeExecute Job 执行前的回调
	// - 注入 Data
	// - 注入 FuncMap
	BeforeExecute(ctx *Context)
	// BeforeOut Job 执行完成，输出前的回调
	// - 可修改输出的内容，如格式化等
	BeforeOut(ctx *Context)
}

var (
	plugins = make(map[string]Plugin)
)

func RegisterPlugin(p Plugin) {
	plugins[p.Name()] = p
}
