package elang

// Alias ...
type Alias struct {
	Name  string
	Value string
}

// NewAlias ...
func NewAlias() *Alias {
	return &Alias{}
}

// String ...
func (a Alias) String() string {
	return "[" + a.Name + ":" + a.Value + "]"
}
