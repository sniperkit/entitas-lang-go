package elang

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseKeyValue(t *testing.T) {
	val := "MyKey:MyValue"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err != nil {
		t.Error(err)
	}
	if k != "MyKey" {
		t.Errorf("test: expected 'MyKey', recieved '%s'", k)
	}
	if v != "MyValue" {
		t.Errorf("test: expected 'MyValue', recieved '%s'", v)
	}
}

func TestParseKeyWithUnderscoreValueWithoutUnderscore(t *testing.T) {
	val := "My_Key:MyValue"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err != nil {
		t.Error(err)
	}
	if k != "My_Key" {
		t.Errorf("test: expected 'MyKey', recieved '%s'", k)
	}
	if v != "MyValue" {
		t.Errorf("test: expected 'MyValue', recieved '%s'", v)
	}
}

func TestParseKeyWithUnderscoreValueWithUnderscore(t *testing.T) {
	val := "My_Key:My_Value"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err != nil {
		t.Error(err)
	}
	if k != "My_Key" {
		t.Errorf("test: expected 'MyKey', recieved '%s'", k)
	}
	if v != "My_Value" {
		t.Errorf("test: expected 'MyValue', recieved '%s'", v)
	}
}

func TestParseKeyWithoutValue(t *testing.T) {
	val := "MyKey"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err != nil {
		t.Error(err)
	}
	if k != "MyKey" {
		t.Errorf("test: expected 'MyKey', recieved '%s'", k)
	}
	if v != "" {
		t.Errorf("test: expected '', recieved '%s'", v)
	}
}

func TestParseKeyWithColonWithoutValue(t *testing.T) {
	val := "MyKey:"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err == fmt.Errorf("Parse identifier failed. Found \"\", expected word") {
		t.Error(err)
	}
	if k != "" {
		t.Errorf("test: expected '', recieved '%s'", k)
	}
	if v != "" {
		t.Errorf("test: expected '', recieved '%s'", v)
	}
}

func TestParseKeyWithNumber(t *testing.T) {
	val := "MyKey0:"
	p := NewParser(strings.NewReader(val))
	k, v, err := p.parseKeyValue()
	if err != nil {
		t.Error(err)
	}
	if k != "MyKey" {
		t.Errorf("test: expected '', recieved '%s'", k)
	}
	if v != "" {
		t.Errorf("test: expected '', recieved '%s'", v)
	}
}
