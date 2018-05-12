package elang

type TargetDecl struct {
	Target        string
	TargetVersion string
}

func NewTargetDecl() *TargetDecl {
	return &TargetDecl{"", "0.0.0"}
}
