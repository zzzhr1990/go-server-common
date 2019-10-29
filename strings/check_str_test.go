package strings_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zzzhr1990/go-server-common/strings"
)

func TestStrString(t *testing.T) {
	Convey("Check", t, func() {
		//
		res := strings.ConvertToCheckString("0ca175b9c0f726a831d895e269332461", "papaya")
		decode, _ := strings.RecoveryCheckString(res)
		So(decode, ShouldEqual, "0ca175b9c0f726a831d895e269332461")
		_, err := strings.RecoveryCheckString(res + "xxxxx")
		So(err, ShouldNotBeNil)
	})
}
