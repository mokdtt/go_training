package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return "(" + b.x.String() + string(b.op) + b.y.String() + ")"
}

func (c call) String() string {
	b := &bytes.Buffer{}
	b.WriteString(c.fn)
	b.WriteString("(")
	sep := ""
	for _, a := range c.args {
		b.WriteString(sep + a.String())
		sep = ","
	}
	b.WriteString(")")
	return b.String()
}
