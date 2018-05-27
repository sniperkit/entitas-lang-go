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

// GetContextWithParameter ...
func (c *ContextDecl) GetContextWithParameter(Parameter string) *Context {
	for _, context := range c.Context {
		if _, ok := context.GetParameter(Parameter); ok {
			return context
		}
	}
	return nil
}

// GetContextWithParameterValue ...
func (c *ContextDecl) GetContextWithParameterValue(Parameter string, Value string) *Context {
	for _, context := range c.Context {
		if value, _ := context.GetParameter(Parameter); value == Value {
			return context
		}
	}
	return nil
}
