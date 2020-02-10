package main

import (
	"fmt"
	"math"
	"testing"
)


type Expr interface {
	Eval(env Env) float64
	// Check 方法报告表达式中的错误，并把表达式中的变量加入 Vars 中
	//Check(vars map[Var]bool) error
}

// Var 表示一个变量，比如 x
type Var string
// literal 是一个数字常量，比如 3.141
type literal float64
// unary 表示一元操作符表达式，比如-x
type unary struct {
	op rune // '+', '-' 中的一个
	x Expr
}
// binary 表示二元操作符表达式，比如 x+y
type binary struct {
	op rune // '+', '-', '*', '/' 中的一个
	x, y Expr
}
// call 表示函数调用表达式，比如 sin(x)
type call struct {
	fn string // one of "pow", "sin", "sqrt" 中的一个
	args []Expr
}

type Env map[Var]float64


func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5/9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5/9 * (F - 32)", Env{"F": 32}, "0"},
		{"5/9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// 仅在表达式变更时才输出
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr := Var(test.expr)
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
