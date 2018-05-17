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
