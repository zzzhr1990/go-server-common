package strings_test

import (
	"testing"

	"unicode/utf8"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zzzhr1990/go-server-common/strings"
)

func TestRandString(t *testing.T) {
	Convey("Rand", t, func() {
		//
		res := strings.RandString(0)
		So(res, ShouldBeEmpty)
		//So(res, ShouldEqual, "9996f2d1dbafecfedd790498b542864dce0a2a7f793a9d930aed5a553ecc37f7")
		So(utf8.RuneCountInString(strings.RandString(128)), ShouldEqual, 128)
		So(strings.RandString(-1), ShouldBeEmpty)
	})
}
