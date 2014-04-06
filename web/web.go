package web

import (
	"log"
	"net/http"
)

func Run() {
	log.Fatal(http.ListenAndServe("/", http.HandlerFunc(router)))
}
