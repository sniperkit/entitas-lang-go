package elang

// ComponentDecl ...
type ComponentDecl struct {
	Name      string
	Parameter KeyValue
	Context   []string
}

// NewComponentDecl ...
func NewComponentDecl() *ComponentDecl {
	return &ComponentDecl{"", make(map[string]string, 0), make([]string, 0)}
}

// AddContext ...
func (c *ComponentDecl) AddContext(Context string) {
	c.Context = append(c.Context, Context)
}
