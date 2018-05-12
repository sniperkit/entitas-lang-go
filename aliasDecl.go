package elang

type AliasDecl struct {
	Alias []*Alias
}

func NewAliasDecl() *AliasDecl {
	return &AliasDecl{make([]*Alias, 0)}
}

func (a *AliasDecl) AddAlias(Alias *Alias) {
	a.Alias = append(a.Alias, Alias)
}
