package tree

type Node interface {
	Eval() error
}

type NodeFunc func() error

func (f NodeFunc) Eval() error { return f() }
