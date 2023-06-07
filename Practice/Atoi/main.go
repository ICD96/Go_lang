package main

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

func atoi(s string) (int, error) {
	n := 0
	e := 0
	j := 0
	for i := len(s); i != 0; i-- {
		e = int(math.Pow(10, float64(j)))
		j++
		t1 := s[:i]
		temp, _ := utf8.DecodeLastRuneInString(t1)
		switch string(temp) {
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
		case "0":
			n = n + 0*e
		default:
			return n, errors.New("Ввели не число")
		}
	}
	return n, nil
}

func main() {
	c, err := atoi("54500")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
