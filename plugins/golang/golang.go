// golang plugin
package golang

import (
	"github.com/v1990/protoc-gen-goku/goku"
)

type goPlugin struct {
	*goku.Generator
}

type ctxKey struct{}

func init() {
	goku.RegisterPlugin(new(goPlugin))
}

func (p *goPlugin) Name() string {
	return "golang"
}

func (p *goPlugin) Init(g *goku.Generator) {
	p.Generator = g

	p.loadPackages(g.Value("GoPackages"))

}

func (p *goPlugin) BeforeExecute(ctx *goku.Context) {
	c := NewContext(ctx)
	ctx.PutValue(ctxKey{}, c)
	ctx.MergeFuncMap(c.FuncMap())
}

func (p *goPlugin) BeforeOut(ctx *goku.Context) {
	f := NewGoFormatter(ctx)
	content := f.Format()
	ctx.SetContent(content)
}

func (p *goPlugin) loadPackages(v interface{}) {
	if v == nil {
		return
	}
	vv, ok := v.(map[string]interface{})
	if !ok {
		return
	}
	for name, line := range vv {
		if pkgPath, ok := line.(string); ok {
			globalPackages[name] = &GoPackage{
				Name:       name,
				ImportPath: pkgPath,
			}
		}
	}
}
