package tree

import (
	"bufio"
	"io"
)

type File struct {
	Dest    string
	Src     io.Reader
	SrcFile string
	SrcFunc func(Context) (io.Reader, error)
}

func (f File) Eval(c Context) error {
	w, err := c.Create(f.Dest)
	if err != nil {
		return err
	}
	defer w.Close()
	r, err := f.getSource(c)
	if err == nil && r != nil {
		_, err = bufio.NewWriter(w).ReadFrom(r)
	}
	return err
}

func (f File) getSource(c Context) (io.Reader, error) {
	if f.Src != nil {
		return f.Src, nil
	}
	if f.SrcFunc != nil {
		return f.SrcFunc(c)
	}
	return c.Open(f.SrcFile)
}

var _ Node = &File{}
