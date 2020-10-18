package goku

import (
	. "github.com/smartystreets/goconvey/convey"
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
