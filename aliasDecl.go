package elang

import "bytes"

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
		if alias.Name == Name {
			return alias
		}
	}
	return nil
}

// GetAliasWithValue ...
func (a *AliasDecl) GetAliasWithValue(Value string) *Alias {
	for _, alias := range a.Alias {
		if alias.Value == Value {
			return alias
		}
	}
	return nil
}

// GetAliasWithNameValue ...
func (a *AliasDecl) GetAliasWithNameValue(Name string, Value string) *Alias {
	for _, alias := range a.Alias {
		if alias.Name == Name && alias.Value == Value {
			return alias
		}
	}
	return nil
}

// String ...
func (a *AliasDecl) String() string {
	var buffer bytes.Buffer
	for i, alias := range a.Alias {
		buffer.WriteString(alias.String())
		if i != len(a.Alias)-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
