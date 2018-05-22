package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var KeyValueListTestData = []struct {
	in    string
	out   map[string]string
	error bool
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"KeyOne,KeyTwo", KeyValue{"KeyOne": "", "KeyTwo": ""}, false},
	/* SET 02 */ {"KeyOne, KeyTwo", KeyValue{"KeyOne": "", "KeyTwo": ""}, false},
	/* SET 03 */ {"KeyOne:Blah, KeyTwo", KeyValue{"KeyOne": "Blah", "KeyTwo": ""}, false},
	/* SET 04 */ {"KeyOne:Blah, KeyTwo:Blah", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}, false},
	/* SET 05 */ {"KeyOne:Blah,KeyTwo:Blah", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}, false},
	/* SET 06 */ {"Key_One:Blah,KeyTwo:Blah", KeyValue{"Key_One": "Blah", "KeyTwo": "Blah"}, false},
	/* SET 07 */ {"Key_One:Blah,_KeyTwo:Blah", KeyValue{"Key_One": "Blah", "_KeyTwo": "Blah"}, false},

	/* UNACCEPTED USAGE. */

	/* SET 08 */ {"Key_One:Blah,_:Blah", KeyValue{}, true},
}

func TestParseKeyValueList(t *testing.T) {
	for i, d := range KeyValueListTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key value list "+d.in, func() {
					kv, err := p.parseKeyValueList()
					Convey(SwapString(d.error, "should return an error", "should not return an error"), func() {
						SwapError(d.error, err)
						for dk, dv := range d.out {
							Convey(SwapString(d.error, "parsed key value list should not contain key", "parsed key value list should contain key "+dk), func() {
								_, ok := kv[dk]
								SwapBoolean(!d.error, ok)
								Convey(SwapString(d.error, "which should contain value "+dv, "with no value"), func() {
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
