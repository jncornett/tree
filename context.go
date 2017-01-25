package tree

import "github.com/spf13/afero"
import "text/template"

type Context interface {
	afero.Fs
	Enter(name string) (Context, error)
	Template(name string) (*template.Template, error)
	Data(key string) (interface{}, error)
}
