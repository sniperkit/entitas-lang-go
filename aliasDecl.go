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

// GetAliasWithName ...
func (a *AliasDecl) GetAliasWithName(Name string) *Alias {
	for _, alias := range a.Alias {
		if alias.AliasName == Name {
			return alias
		}
	}
	return nil
}

// GetAliasWithValue ...
func (a *AliasDecl) GetAliasWithValue(Value string) *Alias {
	for _, alias := range a.Alias {
		if alias.AliasValue == Value {
			return alias
		}
	}
	return nil
}

// GetAliasWithNameValue ...
func (a *AliasDecl) GetAliasWithNameValue(Name string, Value string) *Alias {
	for _, alias := range a.Alias {
		if alias.AliasName == Name && alias.AliasValue == Value {
			return alias
		}
	}
	return nil
}
