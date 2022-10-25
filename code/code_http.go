package code

import (
	"fmt"
	"net/http"
	"sync"
)

type HttpCode struct {
	c          *Code
	httpStatus int
	ext        string
}

var codeWithHttp = sync.Map{}
var UnknownHttpCode = &HttpCode{c: &UnknownCode, httpStatus: http.StatusInternalServerError, ext: "An internal server error occurred"}

func RegisterHttp(coder *Code, httpStatus int, ext string) {
	if coder.Reserved() {
		panic(fmt.Sprintf("coder: %v is reserved as unknownCode error code", coder))
	}

	c := &HttpCode{
		c:          coder,
		httpStatus: httpStatus,
		ext:        ext,
	}

	codeWithHttp.Store(coder.Code(), c)
}

func MustRegisterHttp(coder *Code, httpStatus int, ext string) {
	if coder.Reserved() {
		panic(fmt.Sprintf("coder: %v is reserved as unknownCode error code", coder))
	}

	if _, ok := codeWithHttp.Load(coder.Code()); ok {
		panic(fmt.Sprintf("code: %v already exist", coder))
	}

	c := &HttpCode{
		c:          coder,
		httpStatus: httpStatus,
		ext:        ext,
	}

	codeWithHttp.Store(coder.Code(), c)
}

func (c *HttpCode) HttpStatus() int { return c.httpStatus }

func (c *HttpCode) Ext() string { return c.ext }

func LoadHttpCode(c *Code) *HttpCode {
	if v, ok := codeWithHttp.Load(c.Code()); ok {
		if code, ok := v.(*HttpCode); ok {
			return code
		}

		return UnknownHttpCode
	}

	return UnknownHttpCode
}

func (c *HttpCode) Code() string {
	return c.c.Code()
}
