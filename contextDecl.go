package elang

// ContextDecl ...
type ContextDecl struct {
	Context []*Context
}

// NewContextDecl ...
func NewContextDecl() *ContextDecl {
	return &ContextDecl{make([]*Context, 0)}
}

// AddContext ...
func (c *ContextDecl) AddContext(Context *Context) {
	c.Context = append(c.Context, Context)
}

// GetContextWithName ...
func (c *ContextDecl) GetContextWithName(Name string) *Context {
	for _, context := range c.Context {
		if context.ContextName == Name {
			return context
		}
	}
	return nil
}
