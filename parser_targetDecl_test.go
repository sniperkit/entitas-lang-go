package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var TargetDeclTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"target entitas_csharp", "entitas_csharp"},
	/* SET 02 */ {"target entitascsharp", "entitascsharp"},
	/* SET 03 */ {"target entitas_csharp         ", "entitas_csharp"},
	/* SET 04 */ {"target entitas_csharp       h6uj  ", "entitas_csharp"},
}

var TargetDeclTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {""},
	/* SET 02 */ {"target"},
	/* SET 03 */ {"target _"},
	/* SET 05 */ {"TARGET"},
	/* SET 06 */ {"TARGET _"},
	/* SET 07 */ {"TARGET my_target"},
}

func TestParseTargetDecl(t *testing.T) {
	for i, d := range TargetDeclTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing targetDecl "+d.in, func() {
					targetDecl, err := p.parseTargetDecl()
					Convey("should not return an error ", func() {
						So(err, ShouldBeNil)
						Convey("targetDecl should not be nil ", func() {
							So(targetDecl, ShouldNotBeNil)
							Convey("parsed value should equal "+d.out, func() {
								So(targetDecl.Target, ShouldEqual, d.out)
							})
						})
					})
				})
			})
		})
	}
}

func TestParseErrorTargetDecl(t *testing.T) {
	for i, d := range TargetDeclTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing targetDecl  "+d.in, func() {
					targetDecl, err := p.parseTargetDecl()
					Convey("should return an error ", func() {
						So(err, ShouldNotBeNil)
						Convey("targetDecl should be nil ", func() {
							So(targetDecl, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
