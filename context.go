package elang

type Context struct {
	ContextName      string
	ContextParameter KeyValue
}

func NewContext() *Context {
	return &Context{"", make(KeyValue, 0)}
}

func (c *Context) AddKeyValue(Key string, Value string) {
	c.ContextParameter[Key] = Value
}

func (c *Context) AddKey(Key string) {
	c.ContextParameter[Key] = ""
}
