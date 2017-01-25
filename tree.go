package tree

type Tree []Node

func (tr Tree) Eval(c Context) (err error) {
	for _, node := range tr {
		if err = node.Eval(c); err != nil {
			break
		}
	}
	return
}

var _ Node = Tree{}
