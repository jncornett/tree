package tree

type If struct {
	Func     func(Context) bool
	UseStack bool
	DataKey  string
	Content  Node
}

func (i If) Eval(c Context) (err error) {
	var b bool
	switch {
	case i.Func != nil:
		b = i.Func(c)
	case i.UseStack:
		b = c.Top() == nil
	default:
		b = c.Get(i.DataKey) == nil
	}
	if b {
		err = i.Content.Eval(c)
	}
	return
}
