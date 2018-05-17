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

// AddKeyValue ...
func (c *Context) AddKeyValue(Key string, Value string) {
	c.ContextParameter[Key] = Value
}

// AddKey ...
func (c *Context) AddKey(Key string) {
	c.ContextParameter[Key] = ""
}
