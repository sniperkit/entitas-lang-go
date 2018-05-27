package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var NamespaceTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"my.game.name_space", "my.game.name_space"},
	/* SET 02 */ {"my_game", "my_game"},
	/* SET 03 */ {"my_game         ", "my_game"},
	/* SET 04 */ {"my_game       h6uj  ", "my_game"},
}

var NamespaceTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {""},
	/* SET 02 */ {"my.game."},
	/* SET 03 */ {"_."},
	/* SET 04 */ {"."},
}

func TestParseNamespace(t *testing.T) {
	for i, d := range NamespaceTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing namespace "+d.in, func() {
					str, err := p.parseNamespace()
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

func TestParseErrorNamespace(t *testing.T) {
	for i, d := range NamespaceTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing namespace  "+d.in, func() {
					str, err := p.parseNamespace()
					Convey("should return an error ", func() {
						So(err, ShouldNotBeNil)
						Convey("parsed value should be blank ", func() {
							So(str, ShouldBeBlank)
						})
					})
				})
			})
		})
	}
}
