package web

import (
	"log"
	"net/http"
)

func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(router)))
}
