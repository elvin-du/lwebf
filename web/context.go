package web

import (
	"log"
	"net/http"
)

type parameters map[string]interface{}

type Context struct {
	http.ResponseWriter
	Req    *http.Request
	ip     string
	Params parameters
}

func NewContext(rw http.ResponseWriter, req *http.Request) *Context {
	p := params(req)
	return &Context{rw, req, getIP(req), p}
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

func (this *Context) Abort(status int, body string) {
	this.WriteHeader(status)
	_, err := this.Write([]byte(body))
	if nil != err {
		log.Println(err)
	}
}
