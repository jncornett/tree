package tree

type Dir struct {
	Name     string
	Contents []Node
}

func (d Dir) Eval(c Context) (err error) {
	if c, err = c.Enter(d.Name); err != nil {
		return
	}
	for _, node := range d.Contents {
		if err = node.Eval(c); err != nil {
			break
		}
	}
	return
}

var _ Node = &Dir{}
