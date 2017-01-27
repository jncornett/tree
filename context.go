package tree

import (
	"io"
	"os"
)

type Context interface {
	Create(name string) (io.WriteCloser, error)
	Open(name string) (io.ReadCloser, error)
	Stat(name string) (os.FileInfo, error)
	Mkdir(name string) error
	Enter(name string) (Context, error)
	Push(interface{}) Context
	Top() interface{}
	PushMap(map[string]interface{}) Context
	Get(string) interface{}
}
