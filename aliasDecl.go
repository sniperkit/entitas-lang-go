package elang

// AliasDecl ...
type AliasDecl struct {
	Alias []*Alias
}

// NewAliasDecl ...
func NewAliasDecl() *AliasDecl {
	return &AliasDecl{make([]*Alias, 0)}
}

// AddAlias ...
func (a *AliasDecl) AddAlias(Alias *Alias) {
	a.Alias = append(a.Alias, Alias)
}
