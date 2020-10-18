package goku

import (
	"bytes"
	"fmt"
	"github.com/v1990/protoc-gen-goku/descriptors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func (g *Generator) generateFile(file *descriptors.FileDescriptorProto) {
	defer g.Recover(nil)

	fCtx := g.ctx.WithLoop(LoopFile, file)
	g.executeJobs(fCtx)

	for _, message := range file.GetMessageType() {
		mCtx := fCtx.WithLoop(LoopMessage, message)
		g.executeJobs(mCtx)

		for _, nestedMessage := range message.GetNestedType() {
			nmCtx := mCtx.WithLoop(LoopNestedMessage, nestedMessage)
			g.executeJobs(nmCtx)
		}
		//
		for _, enumObj := range message.GetEnumType() {
			neCtx := mCtx.WithLoop(LoopNestedEnum, enumObj)
			g.executeJobs(neCtx)
		}
	}

	for _, enumObj := range file.GetEnumType() {
		eCtx := fCtx.WithLoop(LoopEnum, enumObj)
		g.executeJobs(eCtx)
	}

	for _, service := range file.GetService() {
		sCtx := fCtx.WithLoop(LoopService, service)
		g.executeJobs(sCtx)

		for _, method := range service.GetMethod() {
			mCtx := sCtx.WithLoop(LoopMethod, method)
			g.executeJobs(mCtx)
		}
	}

}

func (g *Generator) executeJobs(ctx *Context) {
	g.populateCtx(ctx)
	jobs := g.getJobs(ctx)
	//g.Debug("executeJobs: %s -> [%s] %s jobs: %d", ctx.GetFileName(), ctx.Loop(), ctx.object.GetName(), len(jobs))

	for _, job := range jobs {
		//g.Debug(strings.Repeat("*", 40))
		//g.Debug("==== job: %s  [Start]", job.Base)
		//g.Debug(" %+v", job)
		//g.Debug("")

		// 每个job都创建一个新的 ctx
		g.executeJob(job, ctx.withJob(job))

		//g.Debug("")
		//g.Debug("==== job: %s  [Done]", job.Base)
		//g.Debug(strings.Repeat("=", 40))
	}
}

func (g *Generator) executeJob(job Job, ctx *Context) {
	g.populateCtx(ctx)

	// 解析模板
	var tpl *template.Template
	var err error
	if len(job.Template) > 0 {
		tpl, err = template.New("text").Funcs(ctx.tplFuncMap()).Parse(job.Template)
	} else {
		// 先动态解析 TemplatePath
		tplFile := ctx.MustEval(job.TemplatePath)
		tpl, err = template.New(filepath.Base(tplFile)).Funcs(ctx.tplFuncMap()).ParseFiles(tplFile)
	}
	g.FatalOnErr(err, "parse template: job: %s", job.Name)

	// 渲染模板
	writer := bytes.NewBuffer(nil)
	err = tpl.Execute(writer, ctx.Data())
	g.FatalOnErr(err, "execute template. job: %s object: %s", job.Name, ctx.Message().GetName())

	ctx.SetContent(writer.Bytes())

	// 调用插件：处理生成的内容
	ctx.callPlugins(func(plugin Plugin) {
		plugin.BeforeOut(ctx)
	})
	content := ctx.Content()

	// 输出
	if len(job.Out) == 0 || job.Out == "stderr" {
		fmt.Fprintln(os.Stderr, string(content))
	} else {
		//  输出到 g.response.File
		outFileName := ctx.MustEval(job.Out)
		if filepath.IsAbs(outFileName) {
			// 绝对路径: 直接写文件
			err = os.MkdirAll(filepath.Dir(outFileName), 0755)
			g.FatalOnErr(err, "make dir for %s", outFileName)
			err = ioutil.WriteFile(outFileName, content, 0644)
			g.FatalOnErr(err, "write file fail. job:%s file:%s", job.Name, outFileName)
		} else {
			// 相对路径：输出到 g.response.File
			g.addOutFile(outFileName, string(content))
		}

		g.Debug("execute job done: %-30s %-18s %-30s ==> %s",
			job.Name, "["+ctx.Loop()+"]", ctx.Object().GetName(), outFileName)

	}

}

func (g *Generator) getJobs(ctx *Context) []Job {
	jobs := make([]Job, 0)
	for _, job := range g.conf.Jobs {
		if !job.IsEnable(ctx) {
			continue
		}
		jobs = append(jobs, job)
	}
	return jobs
}
