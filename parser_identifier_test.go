package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var IdentifierTestData = []struct {
	in    string
	out   string
	error bool
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"test", "test", false},
	/* SET 02 */ {"test_id", "test_id", false},
	/* SET 03 */ {" test_id", "test_id", false},
	/* SET 04 */ {"      test_id", "test_id", false},
	/* SET 05 */ {"_test_id", "_test_id", false},
	/* SET 06 */ {"_test___ id", "_test___", false},
	/* SET 07 */ {"_t", "_t", false},

	/* UNACCEPTED USAGE. */

	/* SET 08 */ {"_ t", "", true},
	/* SET 09 */ {"_", "", true},
}

func SwapString(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

func SwapError(e bool, err error) {
	if e {
		So(err, ShouldNotBeNil)
	} else {
		So(err, ShouldBeNil)
	}
}

func SwapBoolean(t bool, ok bool) {
	if t {
		So(ok, ShouldBeTrue)
	} else {
		So(ok, ShouldBeFalse)
	}
}

func TestParseIdentifier(t *testing.T) {
	for i, d := range IdentifierTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing identifier "+d.in, func() {
					id, err := p.parseIdentifier()
					Convey(SwapString(d.error, "should return an error", "should not return an error"), func() {
						SwapError(d.error, err)
						Convey("parsed value should equal "+d.out, func() {
							So(id, ShouldEqual, d.out)
						})
					})
				})
			})
		})
	}
}
