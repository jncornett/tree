package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/jncornett/lg"
	"github.com/jncornett/tree"
	"github.com/spf13/afero"
)

type file struct {
	*lg.Logger
	f afero.File
	b bytes.Buffer
	w io.Writer
}

func newFile(f afero.File, l *lg.Logger) *file {
	fl := &file{
		f:      f,
		Logger: l,
	}
	fl.w = io.MultiWriter(&fl.b, fl.f)
	return fl
}

func (f *file) Close() error {
	err := f.f.Close()
	f.Infof("%q", f.b.String())
	return err
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.f.Read(p)
}

func (f *file) ReadAt(p []byte, off int64) (n int, err error) {
	return f.f.ReadAt(p, off)
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	return f.f.Seek(offset, whence)
}

func (f *file) Write(p []byte) (n int, err error) {
	return f.w.Write(p)
}

func (f *file) WriteAt(p []byte, off int64) (n int, err error) {
	panic("not implemented")
}

func (f *file) Name() string {
	panic("not implemented")
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	panic("not implemented")
}

func (f *file) Readdirnames(n int) ([]string, error) {
	panic("not implemented")
}

func (f *file) Stat() (os.FileInfo, error) {
	return f.f.Stat()
}

func (f *file) Sync() error {
	panic("not implemented")
}

func (f *file) Truncate(size int64) error {
	panic("not implemented")
}

func (f *file) WriteString(s string) (ret int, err error) {
	panic("not implemented")
}

type context struct {
	*lg.Logger
	fs   afero.Fs
	data map[string]interface{}
}

func newContext() *context {
	return &context{
		Logger: lg.New(nil, nil, 0),
		fs:     afero.NewMemMapFs(),
		data:   make(map[string]interface{}),
	}
}

func (c *context) Create(name string) (f afero.File, err error) {
	c.Info("Create ", name)
	if f, err = c.fs.Create(name); err != nil {
		return
	}
	f = newFile(f, c.Logger)
	return
}

func (c *context) Mkdir(name string, perm os.FileMode) error {
	c.Info("Mkdir ", name, " ", perm)
	return c.fs.Mkdir(name, perm)
}

func (c *context) MkdirAll(path string, perm os.FileMode) error {
	panic("not implemented")
}

func (c *context) Open(name string) (f afero.File, err error) {
	c.Info("Open ", name)
	if f, err = c.fs.Open(name); err != nil {
		return
	}
	f = newFile(f, c.Logger)
	return
}

func (c *context) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	panic("not implemented")
}

func (c *context) Remove(name string) error {
	panic("not implemented")
}

func (c *context) RemoveAll(path string) error {
	panic("not implemented")
}

func (c *context) Rename(oldname string, newname string) error {
	panic("not implemented")
}

func (c *context) Stat(name string) (os.FileInfo, error) {
	panic("not implemented")
}

func (c *context) Name() string {
	panic("not implemented")
}

func (c *context) Chmod(name string, mode os.FileMode) error {
	panic("not implemented")
}

func (c *context) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("not implemented")
}

func (c *context) Enter(name string) (ctx tree.Context, err error) {
	c.Info("Enter ", name)
	if err = c.Mkdir(name, os.ModePerm); err != nil {
		return
	}
	ctx = &context{
		Logger: c.Logger,
		fs:     afero.NewBasePathFs(c.fs, name),
		data:   c.data,
	}
	return
}

func (c *context) Template(name string) (*template.Template, error) {
	panic("not implemented")
}

func (c *context) Data(key string) (interface{}, error) {
	d, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("%q not found", key)
	}
	return d, nil
}

type project struct{}

const x = `
def main():
    pass

if __name__ == "__main__":
    main()
`

func main() {
	ctx := newContext()
	ctx.data["project"] = &project{}
	tr := tree.Tree{
		&tree.Template{DestFile: "setup.py", SrcTemplFile: "res://setup.py.tmpl", DataKey: "project"},
		&tree.Dir{
			Name: "helloworld",
			Contents: []tree.Node{
				tree.Touch("__init__.py"),
				&tree.Copy{DestFile: "__main__.py", SrcString: x},
			},
		},
	}
	log.Fatal(tr.Eval(ctx))
}
