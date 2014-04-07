package web

import (
	"net/http"
	"reflect"
)

var (
	gControllers = make(map[string]reflect.Value)
)

func Add(i interface{}) {
	key := reflect.TypeOf(i).String()
	val := reflect.ValueOf(i)
	gControllers[key] = val
}

func router(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext(rw, req)
	k := ctx.Params["controler"].(string)
	a := ctx.Params["action"].(string)

	method := gControllers[k+"Controler"].MethodByName(a + "Action")
	p := make([]reflect.Value, 0, 1)
	p[0] = reflect.ValueOf(ctx)
	method.Call(p)
}
