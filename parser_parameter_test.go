package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var ParameterTestData = []struct {
	in  string
	out map[string]string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"(KeyOne,KeyTwo)", KeyValue{"KeyOne": "", "KeyTwo": ""}},
	/* SET 02 */ {"(KeyOne, KeyTwo)", KeyValue{"KeyOne": "", "KeyTwo": ""}},
	/* SET 03 */ {"(KeyOne:Blah, KeyTwo)", KeyValue{"KeyOne": "Blah", "KeyTwo": ""}},
	/* SET 04 */ {"(KeyOne:Blah, KeyTwo:Blah)", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}},
	/* SET 05 */ {"(KeyOne:Blah,KeyTwo:Blah)", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}},
	/* SET 06 */ {"(Key_One:Blah,KeyTwo:Blah)", KeyValue{"Key_One": "Blah", "KeyTwo": "Blah"}},
	/* SET 07 */ {"(Key_One:Blah,_KeyTwo:Blah)", KeyValue{"Key_One": "Blah", "_KeyTwo": "Blah"}},

	/* SYNTAX INVALID BUT OPTIONAL '(' NOT PRESENT SO PARSE PARAMETER RETURNS NIL */
	/* PARAMETER SYNTAX IS OPTIONAL */

	/* SET 08 */ {"Key_One:Blah,_:Blah", nil},
	/* SET 09 */ {"_:,_:Blah", nil},
}

var ParameterErrorTestData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 03 */ {"()"},
	/* SET 04 */ {"(a, )"},
}

func TestParseParameter(t *testing.T) {
	for i, d := range ParameterTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing parameter "+d.in, func() {
					kv, err := p.parseParameter()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						for dk, dv := range d.out {
							Convey("parsed parameter should contain key"+dk, func() {
								_, ok := kv[dk]
								So(ok, ShouldBeTrue)
								Convey("which should contain value "+dv, func() {
									So(kv[dk], ShouldEqual, dv)
								})
							})
						}
					})
				})
			})
		})
	}
}

func TestParseErrorParameter(t *testing.T) {
	for i, d := range ParameterErrorTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing parameter "+d.in, func() {
					kv, err := p.parseParameter()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("key/value list should be nil", func() {
							So(kv, ShouldBeNil)
						})
					})
				})
			})
		})
	}
}
