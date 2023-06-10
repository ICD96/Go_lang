package main

import (
	"fmt"
	"strconv"
)

func sum(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func div(a, b int) int {
	return a / b
}
func mul(a, b int) int {
	return a * b
}

func main() {
	var num [8]int // Вспомнить как сделать массив без границ
	var err error
	var str, temp string
	var glyph [8]string
	j := 0
	str = "12+14"
	result := 0
	for i := 0; i < len(str); i++ {
		_, err = strconv.Atoi(str[i : i+1])
		if err == nil { // Нужно дополнить проверкой на конец строки, иначе он просто не берет оставшееся число
			temp = temp + str[i:i+1]
		} else {
			num[j], _ = strconv.Atoi(temp)
			fmt.Println(num[j])
			glyph[j] = str[i : i+1]
			j++
			temp = ""
		}
	}
	for i := 0; i < j; i++ { // Нужно реализовать чтобы он делал опереции поочередно (сначала умножение потом остально и тд)
		switch glyph[i] {
		case "+":
			result = result + sum(num[i], num[i+1])
		case "-":
			result = result + sub(num[i], num[i+1])
		case "/":
			result = result + div(num[i], num[i+1])
		case "*":
			result = result + mul(num[i], num[i+1])
		}
	}
	fmt.Println(result)
}
