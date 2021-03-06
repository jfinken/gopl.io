// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"bytes"
	"fmt"
)

func (lit literal) String() string {
	return Format(lit)
}
func (v Var) String() string {
	return Format(v)
}
func (a assign) String() string {
	return Format(a)
}
func (u unary) String() string {
	return Format(u)
}
func (b binary) String() string {
	return Format(b)
}
func (c call) String() string {
	return Format(c)
}

// Format formats an expression as a string.
// It does not attempt to remove unnecessary parens.
func Format(e Expr) string {
	var buf bytes.Buffer
	write(&buf, e)
	return buf.String()
}

func write(buf *bytes.Buffer, e Expr) {
	switch e := e.(type) {
	case literal:
		fmt.Fprintf(buf, "%g", e)

	case Var:
		fmt.Fprintf(buf, "%s", e)

	case assign:
		fmt.Fprintf(buf, "(%s=", e.ident)
		write(buf, e.value)
		buf.WriteByte(')')

	case unary:
		fmt.Fprintf(buf, "(%c", e.op)
		write(buf, e.x)
		buf.WriteByte(')')

	case binary:
		buf.WriteByte('(')
		write(buf, e.x)
		fmt.Fprintf(buf, " %c ", e.op)
		write(buf, e.y)
		buf.WriteByte(')')

	case call:
		fmt.Fprintf(buf, "%s(", e.fn)
		for i, arg := range e.args {
			if i > 0 {
				buf.WriteString(", ")
			}
			write(buf, arg)
		}
		buf.WriteByte(')')

	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}
