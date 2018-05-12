package elang

type NamespaceDecl struct {
	Namespace string
}

func NewNamespaceDecl() *NamespaceDecl {
	return &NamespaceDecl{""}
}
