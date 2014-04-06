package web

import (
	"log"
	//. "lwebf/error"
	"net/http"
	"strings"
)

func getIP(req *http.Request) string {
	substr := strings.Split(req.RemoteAddr, ":")
	if 2 != len(substr) {
		log.Println("client ip not found")
		return ""
	}

	return substr[0]
}
