package elang

// Project ...
type Project struct {
	TargetDecl    *TargetDecl
	NamespaceDecl *NamespaceDecl
	ContextDecl   *ContextDecl
	AliasDecl     []*AliasDecl
	ComponentDecl []*ComponentDecl
}

// NewProject ...
func NewProject() *Project {
	return &Project{AliasDecl: make([]*AliasDecl, 0), ComponentDecl: make([]*ComponentDecl, 0)}
}

// AddAliasDecl ...
func (p *Project) AddAliasDecl(AliasDecl *AliasDecl) {
	p.AliasDecl = append(p.AliasDecl, AliasDecl)
}

// AddComponentDecl ...
func (p *Project) AddComponentDecl(ComponentDecl *ComponentDecl) {
	p.ComponentDecl = append(p.ComponentDecl, ComponentDecl)
}
