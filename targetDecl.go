package elang

// TargetDecl ...
type TargetDecl struct {
	Target        string
	TargetVersion string
}

// NewTargetDecl ...
func NewTargetDecl() *TargetDecl {
	return &TargetDecl{"", "0.0.0"}
}
