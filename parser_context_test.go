package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var ContextTestData = []struct {
	in    string
	name  string
	param map[string]string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"my_game(key_one:value_one)", "my_game", KeyValue{"key_one": "value_one"}},
	/* SET 02 */ {"my_game(key_one:value_one, key_two:value_two)", "my_game", KeyValue{"key_one": "value_one", "key_two": "value_two"}},
	/* SET 03 */ {"my_game", "my_game", nil},
}

var ContextTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"my_game()"},
	/* SET 02 */ {"my_game(a, )"},
}

func TestParseContext(t *testing.T) {
	for i, d := range ContextTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing context "+d.in, func() {
					context, err := p.parseContext()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("context should not be nil", func() {
							So(context, ShouldNotBeNil)
							Convey("parsed context name should equal "+d.name, func() {
								So(context.ContextName, ShouldEqual, d.name)
							})
							Convey("parsed context parameter should not be nil", func() {
								So(context.ContextParameter, ShouldNotBeNil)
								for dk, dv := range d.param {
									Convey("context contains parameter with key "+dk, func() {
										value, ok := context.GetParameter(dk)
										Convey("parameter should exist ", func() {
											So(ok, ShouldBeTrue)
											Convey("parameter should contain value "+dv, func() {
												So(value, ShouldEqual, dv)
											})
										})
									})
								}
							})
						})
					})
				})
			})
		})
	}
}

func TestParseErrorContext(t *testing.T) {
	for i, d := range ContextTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing context "+d.in, func() {
					context, err := p.parseContext()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("context should be nil", func() {
							So(context, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
