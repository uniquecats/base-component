package code

import "fmt"

type Code struct {
	service   int
	module    int
	bussiness int
}

var (
	UnknownCode = Code{service: 0, module: 0, bussiness: 0}
)

func (c *Code) Code() string {
	return fmt.Sprintf("%d-%d-%d", c.service, c.module, c.bussiness)
}

func New() *Code { return &Code{} }

func (c *Code) SetModule(module int) *Code {
	if module == 0 {
		panic("module number `0` is reserved.")
	}

	return &Code{
		service:   c.service,
		module:    module,
		bussiness: c.bussiness,
	}
}

func (c *Code) SetBussiness(code int) *Code {
	if code == 0 {
		panic("code number `0` is reserved.")
	}

	return &Code{
		service:   c.service,
		module:    c.module,
		bussiness: code,
	}
}

func (c *Code) SetService(system int) *Code {
	if system == 0 {
		panic("system number `0` is reserved.")
	}

	return &Code{
		service:   system,
		module:    c.module,
		bussiness: c.bussiness,
	}
}

func (c *Code) Reserved() bool { return c.service == 0 || c.module == 0 || c.bussiness == 0 }
