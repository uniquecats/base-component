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

func (c *Code) SetLevel(level int) *Code {
	return &Code{
		level:  level,
		system: c.system,
		module: c.module,
		code:   c.code,
	}
}

func (c *Code) SetModule(module int) *Code {
	if module == 0 {
		panic("moudule number `0` is reserved.")
	}

	return &Code{
		level:  c.level,
		system: c.system,
		module: module,
		code:   c.code,
	}
}

func (c *Code) SetCode(code int) *Code {
	if code == 0 {
		panic("code number `0` is reserved.")
	}

	return &Code{
		level:  c.level,
		system: c.system,
		module: c.module,
		code:   code,
	}
}

func (c *Code) SetSystem(system int) *Code {
	if system == 0 {
		panic("system number `0` is reserved.")
	}

	return &Code{
		level:  c.level,
		system: system,
		module: c.module,
		code:   c.code,
	}
}

func (c *Code) Reserved() bool { return c.system == 0 || c.module == 0 || c.code == 0 }
