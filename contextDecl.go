package elang

type ContextDecl struct {
	Context []*Context
}

func NewContextDecl() *ContextDecl {
	return &ContextDecl{make([]*Context, 0)}
}

func (c *ContextDecl) AddContext(Context *Context) {
	c.Context = append(c.Context, Context)
}
