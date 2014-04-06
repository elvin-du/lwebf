package web

import (
	"net/http"
)

type context struct {
	ip             string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func NewContext(rw http.ResponseWriter, req *http.Request) *context {
	return &context{getIP(req), rw, req}
}

func (this *context) IP() string {
	return this.ip
}
