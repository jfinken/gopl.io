// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

/*
* TODO:
*	- make this the main geth package.
*	- Parse struct, Env is a field.
* 	- implement New(), creates a Parse struct
 */

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) eval(env Env) float64 {
	return env[v]
}

func (l literal) eval(_ Env) float64 {
	return float64(l)
}
func (a assign) eval(env Env) float64 {
	env[a.ident] = float64(a.value)
	return float64(a.value)
}

//!-Eval1

//!+Eval2

// Eval contains the parser to maintain the implicit interface satisfaction
// by the expression types.
type Eval struct {
	parser parser
}

// New constructs an unexported parser object to store the Env as state.
func New() *Eval {
	p := parser{env: Env{}}
	e := Eval{parser: p}
	return &e
}

// Run executes the given string expression returning a float64 result or panic
func (e *Eval) Run(input string) float64 {
	// parse
	expr, err := e.parser.Parse(input)
	if err != nil {
		panic(fmt.Sprintf("unsupported expression: %s. [Error: %s]",
			input, err.Error()))
	}
	// evaluate the expression based on its type
	return expr.eval(e.parser.env)
}
func (u unary) eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.eval(env)
	case '-':
		return -u.x.eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.eval(env) + b.y.eval(env)
	case '-':
		return b.x.eval(env) - b.y.eval(env)
	case '*':
		return b.x.eval(env) * b.y.eval(env)
	case '/':
		return b.x.eval(env) / b.y.eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type fn func(float64) float64

func (c call) eval(env Env) float64 {
	/*	// Map of function name to func
		funcs := map[string]fn{
			"sin": math.Sin,
		}
		return funcs[c.fn](c.args[0].eval(env))
	*/
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].eval(env), c.args[1].eval(env))
	case "sin":
		return math.Sin(c.args[0].eval(env))
	case "cos":
		return math.Cos(c.args[0].eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Eval2
