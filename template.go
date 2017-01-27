package tree

import (
	"io"
	"io/ioutil"
	"text/template"
)

type Template struct {
	Dest                                      string
	SrcTemplFile, SrcTemplString, SrcTemplKey string
	SrcTempl                                  *template.Template
	SrcTemplReader                            io.Reader
	SrcTemplBytes                             []byte
	DataKey                                   string
	DataObject                                interface{}
	UseContextAsData                          bool
	UseStackData                              bool
}

func (t Template) Eval(c Context) (err error) {
	var (
		tmpl *template.Template
		dest io.WriteCloser
	)
	if tmpl, err = t.getTemplate(c); err != nil || tmpl == nil {
		return // TODO add contextual info to error
	}
	data := t.getData(c)
	if dest, err = c.Create(t.Dest); err != nil {
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

func (t Template) getData(c Context) interface{} {
	if t.UseContextAsData {
		return c
	}
	if t.UseStackData {
		return c.Top()
	}
	if t.DataObject != nil {
		return t.DataObject
	}
	return c.Get(t.DataKey)
}

var _ Node = &Template{}
