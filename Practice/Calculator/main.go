package main

import (
	"fmt"
)

type Calculator struct {
	a int
	b int
}

func (c Calculator) Sum(a, b int) int {
	return a + b
}

func (c Calculator) Sub(a, b int) int {
	return a - b
}

func (c Calculator) Div(a, b int) int {
	return a / b
}

func (c Calculator) Mult(a, b int) int {
	return a * b
}

func New(a, b int) Calculator {
	return Calculator{
		a: a,
		b: b,
	}
}

func main() {
	a := New(0, 0)
	c := ""
	fmt.Scan(&a.a, &c, &a.b)
	//fmt.Scan(&a.b)
	//	fmt.Fscan(os.Stdin, &c)
	switch c {
	case "+":
		fmt.Println(a.Sum(a.a, a.b))
	case "-":
		fmt.Println(a.Sub(a.a, a.b))
	case "*":
		fmt.Println(a.Mult(a.a, a.b))
	case "/":
		fmt.Println(a.Div(a.a, a.b))
	}
}
