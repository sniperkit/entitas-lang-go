package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var NamespaceDeclTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"namespace my.game.name_space", "my.game.name_space"},
	/* SET 02 */ {"namespace my_game", "my_game"},
	/* SET 03 */ {"namespace my_game         ", "my_game"},
	/* SET 04 */ {"namespace my_game       h6uj  ", "my_game"},
}

var NamespaceDeclTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {""},
	/* SET 02 */ {"my.game."},
	/* SET 03 */ {"_."},
	/* SET 04 */ {"."},
	/* SET 05 */ {"namespace"},
	/* SET 06 */ {"namespace my.game."},
	/* SET 07 */ {"namespace _."},
	/* SET 08 */ {"namespace ."},
	/* SET 09 */ {"namespface"},
	/* SET 10 */ {"namespgace my.game."},
	/* SET 11 */ {"name3space _."},
	/* SET 12 */ {"namefspace ."},
	/* SET 13 */ {"NAMESPACE"},
	/* SET 14 */ {"NAMESPACE my.game."},
	/* SET 15 */ {"NAMESPACE _."},
	/* SET 16 */ {"NAMESPACE ."},
	/* SET 17 */ {"NAMESPACE my.game"},
	/* SET 18 */ {"NAMESPACE my.game.namespace"},
	/* SET 19 */ {"NAMESPACE my_game.name_space"},
}

func TestParseNamespaceDecl(t *testing.T) {
	for i, d := range NamespaceDeclTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing namespaceDecl "+d.in, func() {
					nsDecl, err := p.parseNamespaceDecl()
					Convey("should not return an error ", func() {
						So(err, ShouldBeNil)
						Convey("namespace should not be nil ", func() {
							So(nsDecl, ShouldNotBeNil)
							Convey("parsed value should equal "+d.out, func() {
								So(nsDecl.Namespace, ShouldEqual, d.out)
							})
						})
					})
				})
			})
		})
	}
}

func TestParseErrorNamespaceDecl(t *testing.T) {
	for i, d := range NamespaceDeclTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing namespaceDecl  "+d.in, func() {
					nsDecl, err := p.parseNamespaceDecl()
					Convey("should return an error ", func() {
						So(err, ShouldNotBeNil)
						Convey("namespace should be nil ", func() {
							So(nsDecl, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
