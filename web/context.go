package web

import (
	"net/http"
)

type parameters map[string]interface{}

type Context struct {
	ip     string
	RW     http.ResponseWriter
	Req    *http.Request
	Params parameters
}

func NewContext(rw http.ResponseWriter, req *http.Request) *Context {
	p := params(req)
	return &Context{getIP(req), rw, req, p}
}

func params(req *http.Request) parameters {
	req.ParseForm()
	if len(req.Form) <= 0 {
		return nil
	}

	p := make(parameters)
	for k, v := range req.Form {
		p[k] = v
	}

	return p
}

func (this *Context) IP() string {
	return this.ip
}
