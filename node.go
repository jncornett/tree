package tree

type Node interface {
	Eval(Context) error
}

type NodeFunc func(Context) error

func (f NodeFunc) Eval(c Context) error { return f(c) }
