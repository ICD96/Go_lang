package main

import "fmt"

func main() {
	var temp []rune
	str := "Привет"
	run := []rune(str)
	fin := ""
	for i := len(str) / 2; i != 0; i-- {
		run = run[:i]
		temp = run[len(run)-1:]
		fin = fin + string(temp)
	}
	fmt.Println(fin)
}
