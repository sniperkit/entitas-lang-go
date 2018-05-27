package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var ContextDeclTestData = []struct {
	in    string
	name  []string
	param []map[string]string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"context my_game", []string{"my_game"}, nil},
	/* SET 02 */ {"context my_game, my_meta", []string{"my_game", "my_meta"}, nil},
	/* SET 03 */ {"context my_game(key_one:value)", []string{"my_game"}, []map[string]string{KeyValue{"key_one": "value"}}},
	/* SET 04 */ {"context my_game(key_one:value), my_meta(key_one_meta:value)", []string{"my_game", "my_meta"}, []map[string]string{KeyValue{"key_one": "value"}, KeyValue{"key_one_meta": "value"}}},
}

var ContextDeclTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"context my_game(), my_meta"},
	/* SET 02 */ {"context my_game(a, ), my_meta"},
	/* SET 03 */ {"CONTEXT my_game(), my_meta"},
	/* SET 04 */ {"CONTEXT my_game(a, ), my_meta"},
	/* SET 05 */ {"CONTEXT my_game, my_meta"},
	/* SET 06 */ {"CONTEXT my_game(key:value), my_meta"},
}

func TestParseContextDecl(t *testing.T) {
	for i, d := range ContextDeclTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing context "+d.in, func() {
					contextDecl, err := p.parseContextDecl()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("contextDecl should not be nil", func() {
							So(contextDecl, ShouldNotBeNil)
							for i, dn := range d.name {
								Convey("parsed contextDecl should contain context with name "+dn, func() {
									context := contextDecl.GetContextWithName(dn)
									Convey("context should not be nil ", func() {
										So(context, ShouldNotBeNil)
									})
									if d.param != nil {
										for dk, dv := range d.param[i] {
											Convey("context contains parameter with key "+dk, func() {
												value, ok := context.GetParameter(dk)
												Convey("parameter should exist", func() {
													So(ok, ShouldBeTrue)
													Convey("parameter should contain value "+dv, func() {
														So(value, ShouldEqual, dv)
													})
												})
											})
										}
									}
								})
							}
						})
					})
				})
			})
		})
	}
}

func TestParseErrorContextDecl(t *testing.T) {
	for i, d := range ContextDeclTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing contextDecl "+d.in, func() {
					contextDecl, err := p.parseContextDecl()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("context should be nil", func() {
							So(contextDecl, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
