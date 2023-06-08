package main

import "fmt"

func main() {
	str := "Привет"
	run := []rune(str)
	fin := ""
	for i := len(str); i != 0; i-- {
		fin = fin + string(run[i-1:i])
	}

	fmt.Println(fin)
}
