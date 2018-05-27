package elang

import (
	"bytes"
)

// Context ...
type Context struct {
	Name      string
	Parameter KeyValue
}

// NewContext ...
func NewContext() *Context {
	return &Context{Parameter: NewKeyValue()}
}

// SetParameterValue ...
func (c *Context) SetParameterValue(Key string, Value string) {
	c.Parameter[Key] = Value
}

// SetParameter ...
func (c *Context) SetParameter(Key string) {
	c.Parameter[Key] = ""
}

// GetParameter ...
func (c *Context) GetParameter(Key string) (value string, ok bool) {
	value, ok = c.Parameter[Key]
	return
}

// String ...
func (c Context) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(c.Name)
	buffer.WriteRune('(')
	buffer.WriteString(c.Parameter.String())
	buffer.WriteRune(')')
	return buffer.String()
}
