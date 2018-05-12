package elang

type Alias struct {
	AliasName  string
	AliasValue string
}

func NewAlias() *Alias {
	return &Alias{"", ""}
}
