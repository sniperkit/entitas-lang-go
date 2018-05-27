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

var KeyValueTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"_a:_"},
	/* SET 02 */ {"_a:"},
	/* SET 03 */ {"_:_"},
	/* SET 04 */ {":"},
	/* SET 05 */ {""},
}

var KeyValueListTestData = []struct {
	in  string
	out map[string]string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"KeyOne,KeyTwo", KeyValue{"KeyOne": "", "KeyTwo": ""}},
	/* SET 02 */ {"KeyOne, KeyTwo", KeyValue{"KeyOne": "", "KeyTwo": ""}},
	/* SET 03 */ {"KeyOne:Blah, KeyTwo", KeyValue{"KeyOne": "Blah", "KeyTwo": ""}},
	/* SET 04 */ {"KeyOne:Blah, KeyTwo:Blah", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}},
	/* SET 05 */ {"KeyOne:Blah,KeyTwo:Blah", KeyValue{"KeyOne": "Blah", "KeyTwo": "Blah"}},
	/* SET 06 */ {"Key_One:Blah,KeyTwo:Blah", KeyValue{"Key_One": "Blah", "KeyTwo": "Blah"}},
	/* SET 07 */ {"Key_One:Blah,_KeyTwo:Blah", KeyValue{"Key_One": "Blah", "_KeyTwo": "Blah"}},
}

var KeyValueListTestErrorData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"Key_One:Blah,_:Blah"},
	/* SET 02 */ {"_:,_:Blah"},
	/* SET 03 */ {"a, "},
}

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

var StringTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"\"\"", ""},
	/* SET 02 */ {"\"System.Collections.Generic<int>\"", "System.Collections.Generic<int>"},
	/* SET 03 */ {"\"int[]\"", "int[]"},
	/* SET 04 */ {"\"jbuglf_178914 &£&*]\"", "jbuglf_178914 &£&*]"},
}

var StringErrorTestData = []struct {
	in string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {""},
	/* SET 02 */ {"_n"},
	/* SET 03 */ {"qkdfqf _n"},
}

var IdentifierTestData = []struct {
	in  string
	out string
}{
	/* ACCEPTED USAGE. */

	/* SET 01 */ {"test", "test"},
	/* SET 02 */ {"test_id", "test_id"},
	/* SET 03 */ {" test_id", "test_id"},
	/* SET 04 */ {"      test_id", "test_id"},
	/* SET 05 */ {"_test_id", "_test_id"},
	/* SET 06 */ {"_test___ id", "_test___"},
	/* SET 07 */ {"_t", "_t"},
}

var IdentifierTestErrorData = []struct {
	in  string
	out string
}{
	/* UNACCEPTED USAGE. */

	/* SET 01 */ {"_ t", ""},
	/* SET 02 */ {"_", ""},
}

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
	for i, d := range KeyValueTestErrorData {
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

func TestParseKeyValueList(t *testing.T) {
	for i, d := range KeyValueListTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key/value list "+d.in, func() {
					kv, err := p.parseKeyValueList()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						for dk, dv := range d.out {
							Convey("parsed key/value list should contain key"+dk, func() {
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

func TestParseErrorKeyValueList(t *testing.T) {
	for i, d := range KeyValueListTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing key value list "+d.in, func() {
					kv, err := p.parseKeyValueList()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("key value list should be nil", func() {
							So(kv, ShouldBeNil)
						})
					})
				})
			})
		})
	}
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

func TestParseString(t *testing.T) {
	for i, d := range StringTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing string  "+d.in, func() {
					str, err := p.parseString()
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

func TestParseErrorString(t *testing.T) {
	for i, d := range StringErrorTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing string  "+d.in, func() {
					str, err := p.parseString()
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

func TestParseIdentifier(t *testing.T) {
	for i, d := range IdentifierTestData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing identifier "+d.in, func() {
					id, err := p.parseIdentifier()
					Convey("should not return an error", func() {
						So(err, ShouldBeNil)
						Convey("parsed value should equal "+d.out, func() {
							So(id, ShouldEqual, d.out)
						})
					})
				})
			})
		})
	}
}

func TestParseErrorIdentifier(t *testing.T) {
	for i, d := range IdentifierTestErrorData {
		Convey("using data set "+strconv.Itoa(i+1), t, func() {
			Convey("when given a new parser", func() {
				p := NewParser(strings.NewReader(d.in))
				Convey("parsing identifier "+d.in, func() {
					id, err := p.parseIdentifier()
					Convey("should return an error", func() {
						So(err, ShouldNotBeNil)
						Convey("parsed value should be blank ", func() {
							So(id, ShouldBeBlank)
						})
					})
				})
			})
		})
	}
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
								So(context.Name, ShouldEqual, d.name)
							})
							Convey("parsed context parameter should not be nil", func() {
								So(context.Parameter, ShouldNotBeNil)
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
											So(alias.Value, ShouldEqual, dv)
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
