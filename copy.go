package tree

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"github.com/spf13/afero"
)

type Copy struct {
	DestFile           string
	SrcString, SrcFile string
	SrcBytes           []byte
	SrcReader          io.Reader
}

func (y Copy) Eval(c Context) (err error) {
	var (
		src  io.ReadCloser
		dest afero.File
	)
	if src, err = y.getSrc(c); err != nil {
		return // TODO add contextual logging, maybe via custom error type?
	}
	defer src.Close()
	if dest, err = c.Create(y.DestFile); err != nil {
		return // TODO add contextual logging, maybe via custom error type?
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	return
}

func (y Copy) getSrc(c Context) (io.ReadCloser, error) {
	switch {
	case y.SrcBytes != nil:
		return ioutil.NopCloser(bytes.NewReader(y.SrcBytes)), nil
	case y.SrcFile != "":
		return c.Open(y.SrcFile)
	}
	return ioutil.NopCloser(strings.NewReader(y.SrcString)), nil
}

var _ Node = &Copy{}

func Touch(dest string) *Copy {
	return &Copy{DestFile: dest}
}
