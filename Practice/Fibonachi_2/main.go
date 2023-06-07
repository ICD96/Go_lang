package main

import "fmt"

func fibonacchi(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacchi(n-1) + fibonacchi(n-2)
}
func main() {
	fmt.Println(fibonacchi(4))
}
