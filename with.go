package tree

type With struct {
	Ctx     func(Context) (Context, error)
	Data    interface{}
	Map     map[string]interface{}
	Content Node
}

func (w With) Eval(c Context) (err error) {
	switch {
	case w.Ctx != nil:
		if c, err = w.Ctx(c); err != nil {
			return
		}
	case w.Data != nil:
		c = c.Push(c)
	case w.Map != nil:
		c = c.PushMap(w.Map)
	}
	return w.Content.Eval(c)
}

var _ Node = &With{}
