package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var AliasTestData = []struct {
	in    string
	name  string
	value string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"id : \"value\"", "id", "value"},
}

var AliasTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"_ : \"value\""},
	/* SET 02 */ {"_a \"value\""},
}

func TestParseAlias(t *testing.T) {
	for i, d := range AliasTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing alias "+d.in, func() {
					alias, err := p.parseAlias()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("alias should not be nil", func() {
							So(alias, ShouldNotBeNil)
							Convey("parsed alias name should equal "+d.name, func() {
								So(alias.Name, ShouldEqual, d.name)
							})
							Convey("parsed alias value should equal "+d.value, func() {
								So(alias.Value, ShouldEqual, d.value)
							})
						})
					})
				})
			})
		})
	}
}

func TestParseErrorAlias(t *testing.T) {
	for i, d := range AliasTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing alias "+d.in, func() {
					alias, err := p.parseAlias()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("alias should be nil", func() {
							So(alias, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
