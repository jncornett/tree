package tree

import (
	"io"
	"io/ioutil"
	"text/template"

	"github.com/spf13/afero"
)

type Template struct {
	DestFile                                  string
	SrcTemplFile, SrcTemplString, SrcTemplKey string
	SrcTempl                                  *template.Template
	SrcTemplReader                            io.Reader
	SrcTemplBytes                             []byte
	DataObject                                interface{}
	DataKey                                   string
}

func (t Template) Eval(c Context) (err error) {
	var (
		tmpl *template.Template
		data interface{}
		dest afero.File
	)
	if tmpl, err = t.getTemplate(c); err != nil {
		return // TODO add contextual info to error
	}
	if data, err = t.getData(c); err != nil {
		return // TODO add contextual info to error
	}
	if dest, err = c.Create(t.DestFile); err != nil {
		return
	}
	defer dest.Close()
	err = tmpl.Execute(dest, data)
	return
}

func (t Template) getTemplate(c Context) (templ *template.Template, err error) {
	var s string
	switch {
	case t.SrcTempl != nil:
		templ = t.SrcTempl
		return // nothing left to do!
	case t.SrcTemplBytes != nil:
		s = string(t.SrcTemplBytes)
	case t.SrcTemplReader != nil:
		var b []byte
		b, err = ioutil.ReadAll(t.SrcTemplReader)
		if err == nil {
			s = string(b)
		}
	case t.SrcTemplKey != "":
		templ, err = c.Template(t.SrcTemplKey)
		return // nothing left to do!
	case t.SrcTemplFile != "":
		var rc io.ReadCloser
		rc, err = c.Open(t.SrcTemplFile)
		if err != nil {
			break
		}
		defer rc.Close()
		var b []byte
		b, err = ioutil.ReadAll(rc)
		if err != nil {
			break
		}
		s = string(b)
	default:
		s = t.SrcTemplString
	}
	if err != nil {
		templ, err = template.New("").Parse(s)
	}
	return
}

func (t Template) getData(c Context) (interface{}, error) {
	if t.DataObject != nil {
		return t, nil
	}
	return c.Data(t.DataKey)
}

var _ Node = &Template{}
