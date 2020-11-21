// golang plugin
package golang

import (
	"bytes"
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
	"strconv"
)

type goPlugin struct {
	gen *goku.Generator
}

func init() {
	goku.RegisterPlugin(new(goPlugin))
}

func (p *goPlugin) Name() string {
	return "golang"
}

func (p *goPlugin) Init(g *goku.Generator) {
	p.gen = g

	p.loadPackages(g.Value("GoPackages"))

}

func (p *goPlugin) BeforeExecute(ctx *goku.Context) {
	c := NewContext(ctx)
	ctx.MergeFuncMap(c.FuncMap())
}

func (p *goPlugin) BeforeOut(ctx *goku.Context) {
	// TODO 考虑采用 go/ast 分析内容，处理import等
	content, err := p.format(ctx.Content())
	if err != nil {
		p.gen.Warn("[golang]format err:%s \n %s \n", err.Error(),
			string(p.lineIndex(content)))
		return
	}

	ctx.SetContent(content)
}

func (p *goPlugin) format(content []byte) ([]byte, error) {
	//content = p.typeCheck(content)

	source, err := format.Source(content)
	return source, err
}

func (p *goPlugin) typeCheck(content []byte) []byte {
	// TODO 清除 unused import
	fset := token.NewFileSet()
	log.Println(1)
	fileAST, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		log.Println("check err:", err)
		os.Exit(1)
	}
	//p.gen.FatalOnErr(err, "parse file err.")
	//ast.SortImports(fset, fileAST)
	// https://github.com/golang/go/issues/23914
	conf := types.Config{
		//IgnoreFuncBodies:         false,
		//FakeImportC:              false,
		//Error: func(err error) {
		//	panic(err)
		//	//log.Println("format:", err.Error())
		//},
		Importer: importer.For("source", nil),
		//Sizes:                    nil,
		//DisableUnusedImportCheck: true,
	}
	log.Println(2)

	files := []*ast.File{fileAST}
	_, err = conf.Check("", fset, files, nil)
	log.Println(2, 1)

	if err != nil {
		log.Println("check err:", err.Error())
		os.Exit(1)
	}

	//if err != nil {
	//	Throws(errors.Errorf("format check: %w", err))
	//}
	//p.gen.FatalOnErr(err, "format check")
	_ = conf
	log.Println(3)

	var buf bytes.Buffer
	printerConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	err = printerConfig.Fprint(&buf, fset, fileAST)
	p.gen.FatalOnErr(err, "generated Go source code could not be reformatted:")
	content = buf.Bytes()

	return content
}

// 为每行代码开头加上行号
func (p *goPlugin) lineIndex(content []byte) []byte {
	i := 0
	body := helper.EachLine(string(content), func(line string) string {
		i++
		return strconv.Itoa(i) + "\t" + line
	})
	return []byte(body)
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
