package userfiles_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	util "github.com/zzzhr1990/go-server-common/userfiles"
)

func TestFormatPath(t *testing.T) {
	Convey("Test Format", t, func() {
		So(util.FormatPath(""), ShouldEqual, "/")
		So(util.FormatPath("/"), ShouldEqual, "/")
		So(util.FormatPath("/\\"), ShouldEqual, "/")
		So(util.FormatPath("\\"), ShouldEqual, "/")
		So(util.FormatPath("/a/.."), ShouldEqual, "/")
		So(util.FormatPath("/../jj"), ShouldEqual, "/jj")
		So(util.FormatPath("/../我爱你"), ShouldEqual, "/我爱你")
		So(util.FormatPath("我爱你"), ShouldEqual, "/我爱你")
		So(util.FormatPath("我\\爱你"), ShouldEqual, "/我/爱你")
		So(util.FormatPath("c/dc/d/c/dc/d/cd/cd/cdc/d/"), ShouldEqual, "/c/dc/d/c/dc/d/cd/cd/cdc/d")
		So(util.FormatPath("////////////"), ShouldEqual, "/")
		So(util.FormatPath("//////..//////"), ShouldEqual, "/")
		So(util.FormatPath("..//////"), ShouldEqual, "/")
		So(util.FormatPath("../abc"), ShouldEqual, "/abc")
		So(util.FormatPath("../abc "), ShouldEqual, "/abc")
		So(util.FormatPath(" xnewu"), ShouldEqual, "/xnewu")
		So(util.FormatPath(" xnewu.j"), ShouldEqual, "/xnewu.j")
		sx, _ := util.GetFormatedIdentity("/abc")
		So(sx, ShouldEqual, "482a7143ac747eff5e5a5992a6016d65")

	})
}
