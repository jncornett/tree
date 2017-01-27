package tree

type Node interface {
	Eval(Context) error
}

type NodeFunc func(Context) error

func (f NodeFunc) Eval(c Context) error { return f(c) }

type NodeList []Node

func (l NodeList) Eval(c Context) (err error) {
	for _, node := range l {
		if err = node.Eval(c); err != nil {
			return
		}
	}
	return
}
