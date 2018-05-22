package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var StringTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"\"\"", ""},
	/* SET 02 */ {"\"System.Collections.Generic<int>\"", "System.Collections.Generic<int>"},
	/* SET 03 */ {"\"int[]\"", "int[]"},
	/* SET 04 */ {"\"jbuglf_178914 &£&*]\"", "jbuglf_178914 &£&*]"},
}

func TestParseString(t *testing.T) {
	for i, d := range StringTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key/value "+d.in, func() {
					str, err := p.parseString()
					Convey("should not return an error ", func() {
						So(err, ShouldBeNil)
						Convey("parsed value should equal "+d.out, func() {
							So(str, ShouldEqual, d.out)
						})
					})
				})
			})
		})
	}
}
