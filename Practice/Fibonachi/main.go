package main

import "fmt"

func fibonacchi(n int) int {
	a, b := 1, 0
	for i := 0; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
func main() {
	fmt.Println(fibonacchi(7))
}
