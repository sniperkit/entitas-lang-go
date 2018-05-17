package elang

// Alias ...
type Alias struct {
	AliasName  string
	AliasValue string
}

// NewAlias ...
func NewAlias() *Alias {
	return &Alias{"", ""}
}
