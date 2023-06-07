package main

import (
	"fmt"
)

// Пока что сделал без структуры, нужно будет уточнить правильно ли это
type Calculator struct {
}

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}

func Mult(a, b int) int {
	return a * b
}

func New() Calculator {
	return Calculator{}
}

func main() {
	a := 0
	b := 0
	c := ""
	fmt.Scan()
	switch c {
	case "+":
		fmt.Println(Sum(a, b))
	case "-":
		fmt.Println(Sub(a, b))
	case "*":
		fmt.Println(Mult(a, b))
	case "/":
		fmt.Println(Div(a, b))
	}
}
