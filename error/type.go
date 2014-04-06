package error

import (
	"encoding/json"
	"log"
)

type Error struct {
	Code uint64
	Msg  string
}

func NewError(code uint64, msg string) *Error {
	return &Error{code, msg}
}

func (this *Error) Error() string {
	return this.Msg
}

func (this *Error) String() string {
	bin, err := json.Marshal(this)
	if nil != err {
		log.Println(err)
		return ""
	}
	return string(bin)
}
