package web

import (
	"net/http"
)

type parameters map[string]interface{}

type context struct {
	ip             string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         parameters
}

func NewContext(rw http.ResponseWriter, req *http.Request) *context {
	p := params(req)
	return &context{getIP(req), rw, req, p}
}

func params(req *http.Request) parameters {
	req.ParseForm()
	if len(req.Form) > 0 {
		p := make(parameters)
		for k, v := range req.Form {
			p[k] = v
		}
		return p
	}

	return nil
}

func (this *context) IP() string {
	return this.ip
}
