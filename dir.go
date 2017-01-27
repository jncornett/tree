package tree

type Dir struct {
	Name    string
	Content Node
}

func (d Dir) Eval(c Context) (err error) {
	if c, err = c.Enter(d.Name); err != nil {
		return
	}
	if d.Content != nil {
		err = d.Content.Eval(c)
	}
	return
}
