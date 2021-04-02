// golang plugin
package golang

import (
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"reflect"
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

	g.GetGlobalCtx().MergeFuncMap(goku.FuncMap{
		"RegisterGoPackage": func(pkg string) *GoPackage {
			return RegisterPackage(pkg, "")
		},
	})

	p.loadPackages(g.Value("GoPackages"))
	//p.Debug("GoPackages: %#v", g.Value("GoPackages"))
	p.Debug("GoPackages: %s", helper.ShowJSON(globalPackages, 1))
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

	pv := reflect.Indirect(reflect.ValueOf(v))
	if pv.Kind() != reflect.Map {
		p.Warn("GoPackages NOT Map. got: %s", reflect.Indirect(pv).Type())
		return
	}
	keys := pv.MapKeys()
	for _, k := range keys {
		val := pv.MapIndex(k)

		name := k.Interface().(string)
		pkgPath := val.Interface().(string)

		globalPackages[name] = &GoPackage{
			Name:       name,
			ImportPath: pkgPath,
		}
	}
}
