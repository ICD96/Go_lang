package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello"
	s1 := ""
	t := ""
	for i := len(s); i != 0; i-- {
		t = s[:i]
		temp, _ := utf8.DecodeLastRuneInString(t)
		if temp != 65533 {
			s1 = s1 + string(temp)
		}
	}
	fmt.Println(s1)
}
