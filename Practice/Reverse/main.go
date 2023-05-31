package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello"
	rev := ""
	t := ""
	for i := len(str); i != 0; i-- {
		t = str[:i]
		t1, _ := utf8.DecodeLastRuneInString(t)
		if t1 != 65533 {
			rev = rev + string(t1)
		}
	}
	fmt.Println(rev)
}
