package main

import (
	"fmt"
	"math"
)

func atoi(s string) int {
	n := 0
	e := 0
	for i := 1; i <= len(s); i++ {
		e = int(math.Pow(10, float64(i)))
		fmt.Println(s[i:]) // Тут хуйня не работает
		switch s[i:] {
		case "1":
			n = n + 1*e
		case "2":
			n = n + 2*e
		case "3":
			n = n + 3*e
		case "4":
			n = n + 4*e
		case "5":
			n = n + 5*e
		case "6":
			n = n + 6*e
		case "7":
			n = n + 7*e
		case "8":
			n = n + 8*e
		case "9":
			n = n + 9*e
		}
	}
	return n
}

func main() {
	c := atoi("14")
	fmt.Println(c)
}
