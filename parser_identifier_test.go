package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var IdentifierTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"test", "test"},
	/* SET 02 */ {"test_id", "test_id"},
	/* SET 03 */ {" test_id", "test_id"},
	/* SET 04 */ {"      test_id", "test_id"},
	/* SET 05 */ {"_test_id", "_test_id"},
	/* SET 06 */ {"_test___ id", "_test___"},
	/* SET 07 */ {"_t", "_t"},
}

var IdentifierTestErrorData = []struct {
	in  string
	out string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"_ t", ""},
	/* SET 02 */ {"_", ""},
}

func TestParseIdentifier(t *testing.T) {
	for i, d := range IdentifierTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing identifier "+d.in, func() {
					id, err := p.parseIdentifier()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("parsed value should equal "+d.out, func() {
							So(id, ShouldEqual, d.out)
						})
					})
				})
			})
		})
	}
}

func TestParseErrorIdentifier(t *testing.T) {
	for i, d := range IdentifierTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing identifier "+d.in, func() {
					id, err := p.parseIdentifier()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("parsed value should be blank ", func() {
							So(id, ShouldBeBlank)
						})
					})
				})
			})
		})
	}
}
