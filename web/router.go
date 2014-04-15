package web

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var (
	E_CONTROLLER_NOT_FOUND = errors.New("controller not found")
)

var (
	gControllers = make(map[string]reflect.Value)
)

func Add(i interface{}) {
	key := reflect.TypeOf(i).String()
	val := reflect.ValueOf(i)
	gControllers[key] = val
}

var (
	REG_URL = regexp.MustCompile(`^/(\w+)(/?)(\w*)`)
)

func router(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext(rw, req)

	pattern := REG_URL.FindStringSubmatch(req.URL.Path)
	c, a := pattern[1], pattern[3]
	if a == "" {
		a = "Index"
	}

	m, err := LookupController(c, a)
	if nil != err {
		log.Println(err)
		return
	}

	m.Call([]reflect.Value{reflect.ValueOf(ctx)})
}

func LookupController(c, a string) (*reflect.Value, error) {
	c2 := strings.ToUpper(c[:1]) + strings.ToLower(c[1:]) + "Controller"
	a2 := strings.ToUpper(a[:1]) + strings.ToLower(a[1:]) + "Action"

	if cv, ok := gControllers[c2]; ok {
		if m := cv.MethodByName(a2); m.IsValid() {
			return &m, nil
		}
	}

	return nil, E_CONTROLLER_NOT_FOUND
}
