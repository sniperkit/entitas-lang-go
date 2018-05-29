package elang

// ComponentDecl ...
type ComponentDecl struct {
	Name      string
	Parameter KeyValue
	Context   []*Context
}

// NewComponentDecl ...
func NewComponentDecl() *ComponentDecl {
	return &ComponentDecl{Parameter: make(map[string]string, 0), Context: make([]*Context, 0)}
}

// AddContext ...
func (c *ComponentDecl) AddContext(Context *Context) {
	c.Context = append(c.Context, Context)
}

// SetParameterValue ...
func (c *ComponentDecl) SetParameterValue(Key string, Value string) {
	c.Parameter[Key] = Value
}

// SetParameter ...
func (c *ComponentDecl) SetParameter(Key string) {
	c.Parameter[Key] = ""
}

// GetParameter ...
func (c *ComponentDecl) GetParameter(Key string) (value string, ok bool) {
	value, ok = c.Parameter[Key]
	return
}
