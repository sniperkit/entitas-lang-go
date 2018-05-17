package elang

import (
	"strings"
	"testing"
)

func TestParseEmptyString(t *testing.T) {
	val := `""`
	p := NewParser(strings.NewReader(val))
	str, err := p.parseString()
	if err != nil {
		t.Error(err)
	}
	if str != "" {
		t.Errorf("test: expected '', recieved '%s'", str)
	}
}

func TestParseStringWithIllegal(t *testing.T) {
	val := `"&*(<>[]"`
	p := NewParser(strings.NewReader(val))
	str, err := p.parseString()
	if err != nil {
		t.Error(err)
	}
	if str != "&*(<>[]" {
		t.Errorf("test: expected '&*(<>[]', recieved '%s'", str)
	}
}
