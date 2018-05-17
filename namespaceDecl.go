package elang

// NamespaceDecl ...
type NamespaceDecl struct {
	Namespace string
}

// NewNamespaceDecl ...
func NewNamespaceDecl() *NamespaceDecl {
	return &NamespaceDecl{""}
}
