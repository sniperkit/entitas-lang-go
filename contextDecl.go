package elang

import "bytes"

// ContextDecl ...
type ContextDecl struct {
	Context []*Context
}

// NewContextDecl ...
func NewContextDecl() *ContextDecl {
	return &ContextDecl{make([]*Context, 0)}
}

// AddContext ...
func (ContextDecl *ContextDecl) AddContext(Context *Context) {
	ContextDecl.Context = append(ContextDecl.Context, Context)
}

// GetContextWithName ...
func (ContextDecl *ContextDecl) GetContextWithName(Name string) *Context {
	for _, Context := range ContextDecl.Context {
		if Context.Name == Name {
			return Context
		}
	}
	return nil
}

// GetContextSliceWithName ...
func (ContextDecl *ContextDecl) GetContextSliceWithName(Name ...string) []*Context {
	ContextSlice := make([]*Context, 0)
	for _, Context := range ContextDecl.Context {
		for _, Name := range Name {
			if Context.Name == Name {
				ContextSlice = append(ContextSlice, Context)
			}
		}

	}
	return ContextSlice
}

// GetContextWithParameter ...
func (ContextDecl *ContextDecl) GetContextWithParameter(Parameter string) *Context {
	for _, Context := range ContextDecl.Context {
		if _, Ok := Context.GetParameter(Parameter); Ok {
			return Context
		}
	}
	return nil
}

// GetContextWithParameterValue ...
func (ContextDecl *ContextDecl) GetContextWithParameterValue(Parameter string, Value string) *Context {
	for _, Context := range ContextDecl.Context {
		if Value, _ := Context.GetParameter(Parameter); Value == Value {
			return Context
		}
	}
	return nil
}

// String ...
func (ContextDecl ContextDecl) String() string {
	var buffer bytes.Buffer
	for i, context := range ContextDecl.Context {
		buffer.WriteString(context.String())
		if i != len(ContextDecl.Context)-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
