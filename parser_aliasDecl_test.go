package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var AliasDeclTestData = []struct {
	in string
	nv map[string]string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"alias id : \"value\"", KeyValue{"id": "value"}},
	/* SET 01 */ {"alias id : \"value\" id_two : \"another value\"", KeyValue{"id": "value", "id_two": "another value"}},
}

var AliasDeclTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"alias _ : \"value\""},
	/* SET 02 */ {"alias _a \"value\""},
	/* SET 03 */ {"ALIAS _ : \"value\""},
	/* SET 04 */ {"ALIAS _a \"value\""},
	/* SET 06 */ {"ALIAS id : \"value\""},
}

func TestParseAliasDecl(t *testing.T) {
	for i, d := range AliasDeclTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing aliasDecl "+d.in, func() {
					aliasDecl, err := p.parseAliasDecl()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("aliasDecl should not be nil", func() {
							So(aliasDecl, ShouldNotBeNil)
							for dk, dv := range d.nv {
								Convey("parsed AliasDecl contains alias with name "+dk, func() {
									alias := aliasDecl.GetAliasWithName(dk)
									Convey("alias should not be nil ", func() {
										So(alias, ShouldNotBeNil)
										Convey("alias should contain value "+dv, func() {
											So(alias.AliasValue, ShouldEqual, dv)
										})
									})
								})
							}
						})
					})
				})
			})
		})
	}
}

func TestParseErrorAliasDecl(t *testing.T) {
	for i, d := range AliasDeclTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing aliasDecl "+d.in, func() {
					aliasDecl, err := p.parseAliasDecl()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("AliasDecl should be nil", func() {
							So(aliasDecl, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
