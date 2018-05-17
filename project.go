package elang

// Project ...
type Project struct {
	TargetDecl    *TargetDecl
	NamespaceDecl *NamespaceDecl
	ContextDecl   *ContextDecl
	AliasDecl     []*AliasDecl
}

// NewProject ...
func NewProject() *Project {
	return &Project{AliasDecl: make([]*AliasDecl, 0)}
}

// AddAliasDecl ...
func (p *Project) AddAliasDecl(AliasDecl *AliasDecl) {
	p.AliasDecl = append(p.AliasDecl, AliasDecl)
}
