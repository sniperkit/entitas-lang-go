package elang

// Context ...
type Context struct {
	ContextName      string
	ContextParameter KeyValue
}

// NewContext ...
func NewContext() *Context {
	return &Context{"", make(KeyValue, 0)}
}

// SetParameterValue ...
func (c *Context) SetParameterValue(Key string, Value string) {
	c.ContextParameter[Key] = Value
}

// SetParameter ...
func (c *Context) SetParameter(Key string) {
	c.ContextParameter[Key] = ""
}

// GetParameter ...
func (c *Context) GetParameter(Key string) (value string, ok bool) {
	value, ok = c.ContextParameter[Key]
	return
}
