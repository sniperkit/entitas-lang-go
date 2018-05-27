package elang

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var KeyValueTestData = []struct {
	in   string
	kout string
	vout string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"MyKey:MyValue", "MyKey", "MyValue"},
	/* SET 02 */ {"MyKey : MyValue", "MyKey", "MyValue"},
	/* SET 03 */ {"MyKey: MyValue", "MyKey", "MyValue"},
	/* SET 04 */ {"MyKey: MyValue", "MyKey", "MyValue"},
	/* SET 05 */ {"MyKey :MyValue", "MyKey", "MyValue"},
	/* SET 06 */ {"MyKey :MyValue", "MyKey", "MyValue"},
	/* SET 07 */ {"My_Key:MyValue", "My_Key", "MyValue"},
	/* SET 08 */ {"My_Key : MyValue", "My_Key", "MyValue"},
	/* SET 09 */ {"My_Key: MyValue", "My_Key", "MyValue"},
	/* SET 10 */ {"My_Key: MyValue", "My_Key", "MyValue"},
	/* SET 09 */ {"My_Key :MyValue", "My_Key", "MyValue"},
	/* SET 10 */ {"My_Key :MyValue", "My_Key", "MyValue"},
	/* SET 11 */ {"MyKey:My_Value", "MyKey", "My_Value"},
	/* SET 12 */ {"MyKey : My_Value", "MyKey", "My_Value"},
	/* SET 13 */ {"MyKey: My_Value", "MyKey", "My_Value"},
	/* SET 14 */ {"MyKey: My_Value", "MyKey", "My_Value"},
	/* SET 15 */ {"MyKey :My_Value", "MyKey", "My_Value"},
	/* SET 16 */ {"MyKey :My_Value", "MyKey", "My_Value"},
	/* SET 17 */ {"My_Key:My_Value", "My_Key", "My_Value"},
	/* SET 18 */ {"My_Key : My_Value", "My_Key", "My_Value"},
	/* SET 19 */ {"My_Key: My_Value", "My_Key", "My_Value"},
	/* SET 20 */ {"My_Key: My_Value", "My_Key", "My_Value"},
	/* SET 21 */ {"My_Key :My_Value", "My_Key", "My_Value"},
	/* SET 22 */ {"My_Key :My_Value", "My_Key", "My_Value"},
	/* SET 23 */ {"MyKey", "MyKey", ""},
	/* SET 24 */ {"My_Key", "My_Key", ""},
}

var KeyValueErrorTestData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"_a:_"},
	/* SET 02 */ {"_a:"},
	/* SET 03 */ {"_:_"},
	/* SET 04 */ {":"},
	/* SET 05 */ {""},
}

func TestParseKeyValue(t *testing.T) {
	for i, d := range KeyValueTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key/value "+d.in, func() {
					k, v, err := p.parseKeyValue()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("parsed key should equal "+d.kout, func() {
							So(k, ShouldEqual, d.kout)
						})
						Convey("parsed value should equal "+d.vout, func() {
							So(v, ShouldEqual, d.vout)
						})
					})
				})
			})
		})
	}
}

func TestParseErrorKeyValue(t *testing.T) {
	for i, d := range KeyValueErrorTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key/value "+d.in, func() {
					k, v, err := p.parseKeyValue()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("parsed key should equal be blank ", func() {
							So(k, ShouldBeBlank)
						})
						Convey("parsed value should equal be blank", func() {
							So(v, ShouldBeBlank)
						})
					})
				})
			})
		})
	}
}
