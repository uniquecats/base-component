package code

import "fmt"

type Code struct {
	level  int
	system int
	module int
	code   int
}

var (
	UnknownCode = Code{level: 0, system: 0, module: 0, code: 0}
)

func (c *Code) Code() string {
	return fmt.Sprintf("%d-%d-%d-%d", c.level, c.system, c.module, c.code)
}

func New() *Code { return &Code{} }

func (c *Code) SetLevel(level int) { c.level = level }

func (c *Code) SetModule(module int) {
	if module == 0 {
		panic("moudule number `0` is reserved.")
	}

	c.module = module
}

func (c *Code) SetCode(code int) {
	if code == 0 {
		panic("code number `0` is reserved.")
	}

	c.code = code
}

func (c *Code) SetSystem(system int) {
	if system == 0 {
		panic("system number `0` is reserved.")
	}

	c.system = system
}

func (c *Code) Reserved() bool { return c.system == 0 || c.module == 0 || c.code == 0 }

func (c *Code) Copy() *Code {
	return &Code{
		level:  c.level,
		system: c.system,
		module: c.module,
		code:   c.code,
	}
}
