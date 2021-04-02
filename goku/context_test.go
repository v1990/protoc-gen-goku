package goku

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"strings"
	"testing"
)

func Test_Functions(t *testing.T) {
	t.Log(strings.SplitAfter("google.golang.org/protobuf/types/descriptorpb", "/"))

}

func TestContext_Content(t *testing.T) {
	Convey("TestContext_Content", t, func() {

		ctx1 := Context{}

		ctx1.SetContent([]byte("abc"))
		So(string(ctx1.Content()), ShouldEqual, "abc")

		ctx2 := ctx1.copy()
		ctx2.SetContent([]byte("bbb"))
		So(string(ctx1.Content()), ShouldEqual, "abc")
		So(string(ctx2.Content()), ShouldEqual, "bbb")

	})

}

// 测试：宏替换
func Test_MacroReplace(t *testing.T) {

	macros := map[string]string{
		"K1": "hello",
		"K2": "world",

		"M1": "{{#K1}} {{#K2}}!",
		"M2": "{{#M1}} {{#K1}} {{#K2}}",
	}
	ret := MacroReplace("{{#M1}} {{#M2}} {{#M4}}", macros)
	t.Log(ret)

}

var macrosReg = regexp.MustCompile(`\{\{#([a-zA-Z][a-zA-Z0-9_]*)\}\}`)

func MacroReplace(src string, macros map[string]string) string {
	for macrosReg.MatchString(src) {
		src = macrosReg.ReplaceAllStringFunc(src, func(s string) string {
			sm := macrosReg.FindStringSubmatch(s)
			if len(sm) < 1 {
				panic(fmt.Errorf("internal error: invalid macrosReg"))
			}
			if r, ok := macros[sm[1]]; ok {
				return r
			}
			return ""
		})
	}
	return src
}
