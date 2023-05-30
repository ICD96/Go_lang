package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет"
	s1 := ""
	t := ""
	for i := len(s); i != 0; i = i - 2 {
		t = s[:i]
		temp, _ := utf8.DecodeLastRuneInString(t)
		s1 = s1 + string(temp)
	}
	fmt.Println(s1)
}
