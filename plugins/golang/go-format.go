package golang

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/v1990/protoc-gen-goku/goku"
	"github.com/v1990/protoc-gen-goku/helper"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
	"log"
	"strconv"
)

type GoFormatter struct {
	*goku.Context
	markedImports map[string]*GoPackage
	buf           *bytes.Buffer
}

func NewGoFormatter(ctx *goku.Context) *GoFormatter {
	c := ctx.Value(ctxKey{}).(*Context)

	return &GoFormatter{
		Context:       ctx,
		markedImports: c.markedImports,
		buf:           bytes.NewBuffer(ctx.Content()),
	}
}

func (c *GoFormatter) Format() []byte {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[golang]format err:%s \n %s \n", err,
				string(c.lineIndex()))
			panic(err)
		}
	}()

	c.addMarkedImports()
	c.cleanImports()
	c.format()

	return c.buf.Bytes()
}

func (c *GoFormatter) format() {
	content, err := format.Source(c.buf.Bytes())
	c.ThrowsOnErr(errors.Wrap(err, "format source err."))
	c.setContent(content)

}

// clean unused import
//  https://stackoverflow.com/questions/55645258/remove-unused-imports-with-ast-package
func (c *GoFormatter) cleanImports() {
	content, err := imports.Process("", c.buf.Bytes(), nil)
	c.ThrowsOnErr(errors.Wrap(err, "clean unused import error."))
	c.setContent(content)
}
func (c *GoFormatter) setContent(content []byte) {
	c.buf.Reset()
	c.buf.Write(content)
}
func (c *GoFormatter) addMarkedImports() {
	if len(c.markedImports) == 0 {
		return
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", c.buf.Bytes(), parser.AllErrors)
	c.ThrowsOnErr(errors.Wrap(err, "ParseFile."))

	for _, p := range c.markedImports {
		added := astutil.AddNamedImport(fset, f, p.Name, p.ImportPath)
		if added {
			c.Debug("add import: %s %s ", p.Name, p.ImportPath)
		}
	}

	c.printAstFile(fset, f)
}

func (c *GoFormatter) printAstFile(fset *token.FileSet, f *ast.File) {
	c.buf.Reset()
	printerConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	err := printerConfig.Fprint(c.buf, fset, f)
	c.ThrowsOnErr(errors.Wrap(err, "generated Go source code could not be reformatted"))
}

//
//func (c *GoFormatter) typeCheck(content []byte) []byte {
//	// TODO 清除 unused import
//	fset := token.NewFileSet()
//	log.Println(1)
//	f, err := parser.ParseFile(fset, "", content, parser.AllErrors)
//	if err != nil {
//		log.Println("check err:", err)
//		os.Exit(1)
//	}
//	// info.Uses allows to lookup import paths for identifiers.
//	info := &types.Info{
//		Uses: make(map[*ast.Ident]types.Object),
//	}
//	//c.Generator.FatalOnErr(err, "parse file err.")
//	//ast.SortImports(fset, f)
//	// https://github.com/golang/go/issues/23914
//	conf := types.Config{
//		//IgnoreFuncBodies:         false,
//		//FakeImportC:              false,
//		//Error: func(err error) {
//		//	panic(err)
//		//	//log.Println("format:", err.Error())
//		//},
//		//Importer: importer.Default(),
//		Importer: importer.For("source", nil),
//		//Sizes:                    nil,
//		//DisableUnusedImportCheck: true,
//	}
//	log.Println(2)
//
//	files := []*ast.File{f}
//	_, err = conf.Check("", fset, files, info)
//	log.Println(2, 1)
//
//	if err != nil {
//		log.Println("check err:", err.Error())
//		os.Exit(1)
//	}
//
//	//if err != nil {
//	//	ThrowsOnErr(errors.Errorf("format check: %w", err))
//	//}
//	//c.Generator.FatalOnErr(err, "format check")
//	_ = conf
//	log.Println(3)
//
//	var buf bytes.Buffer
//	printerConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
//	err = printerConfig.Fprint(&buf, fset, f)
//	c.FatalOnErr(err, "generated Go source code could not be reformatted:")
//	content = buf.Bytes()
//
//	return content
//}

// 为每行代码开头加上行号
func (c *GoFormatter) lineIndex() []byte {
	i := 0
	body := helper.EachLine(string(c.buf.String()), func(line string) string {
		i++
		return strconv.Itoa(i) + "\t" + line
	})
	return []byte(body)
}
