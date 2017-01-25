package tree_test

import "github.com/jncornett/tree"

func ExampleTree() {
	tr := tree.Tree{
		&tree.Template{DestFile: "setup.py", SrcTemplFile: "res://setup.py.tmpl", Key: "project"},
		&tree.Dir{
			Name: "helloworld",
			Contents: []tree.Node{
				tree.Touch("__init__.py"),
				tree.Touch("__main__.py"),
			},
		},
	}
	tr.Eval(nil)
	// tr.Eval(ctx)
}
