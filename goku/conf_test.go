package goku

//
//import (
//	"fmt"
//	. "github.com/smartystreets/goconvey/convey"
//	"runtime"
//	"testing"
//)
//
//func TestCondition_ifOK(t *testing.T) {
//	engine := new(Generator)
//	ctx := newContext(engine)
//	ctx.MergeData(globalData)
//	ctx.MergeFuncMap(globalFuncMap)
//	ctx.MergeData(Data{
//		"k1": true,
//		"k2": false,
//		"k3": "v3",
//		"k4": "true",
//	})
//
//	_, file, _, _ := runtime.Caller(0)
//	t.Log(file)
//
//	tests := []struct {
//		If   string
//		ifOK bool
//	}{
//		{"true", true},
//		{"{{.k1}}", true},
//		{"{{.k2}}", false},
//		{"{{.k4}}", true},
//		{"{{eq .k2 false}}", true},
//		{`{{eq .k3  "v3"}}`, true},
//		{`{{in .k3  "v1" "v3"}}`, true},
//		{`{{notIn .k3  "v1" "v3"}}`, false},
//		// 当前文件：肯定存在
//		{fmt.Sprintf(`{{fileExists "%s"}}`, file), true},
//		// 不存在的文件
//		{fmt.Sprintf(`{{fileExists "%s"}}`, "/NotExists/"+file), false},
//
//		// 注意： 不同类型不能比较
//		// err:template: text:1:2: executing "text" at <eq .k2 "false">:
//		// error calling eq: incompatible types for comparison
//		//{`{{eq .k2 "false"}}`, true},
//	}
//
//	Convey("TestCondition_ifOK", t, func() {
//		for _, tt := range tests {
//			Convey(tt.If, func() {
//				cond := Enable{If: tt.If}
//				So(cond.ifOK(ctx), ShouldEqual, tt.ifOK)
//			})
//		}
//	})
//
//}
//
//func TestJob_getOutPlugins(t *testing.T) {
//	Convey("", t, func() {
//		tests := []struct {
//			Plugins    []string
//			OutPlugins []string
//			except     []string
//		}{
//			{
//				[]string{"p1", "p2", "p3"},
//				[]string{"-p1", "p4", "p3"},
//				[]string{"p4", "p3", "p2"},
//			},
//			{
//				[]string{"p1"},
//				[]string{"-p1"},
//				[]string{},
//			},
//		}
//		for _, tt := range tests {
//			//job := Job{
//			//	Plugins:    tt.Plugins,
//			//}
//			//So(job.getOutPlugins(), ShouldResemble, tt.except)
//		}
//
//	})
//}
//
//func TestCondition_OK(t *testing.T) {
//	type fields struct {
//		Loop []string
//		If   IfCondition
//	}
//	type args struct {
//		e   *Generator
//		ctx *Context
//	}
//	gen := new(Generator)
//	fCtx := newContext(gen).WithLoop(LoopFile, nil)
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		{
//			name:   "a",
//			fields: fields{Loop: []string{}},
//			args:   args{e: gen, ctx: fCtx},
//			want:   true,
//		},
//		{
//			name:   "b",
//			fields: fields{Loop: []string{"message", "enum"}},
//			args:   args{e: gen, ctx: fCtx.WithLoop(LoopMessage, nil)},
//			want:   true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			//c := Enable{
//			//	Loop: tt.fields.Loop,
//			//	If:   tt.fields.If,
//			//}
//			//if got := c.OK(tt.args.ctx); got != tt.want {
//			//	t.Errorf("OK() = %v, want %v", got, tt.want)
//			//}
//		})
//	}
//}
