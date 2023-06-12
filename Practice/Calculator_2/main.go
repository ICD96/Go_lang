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
	str = "1+2*3"
	result := 0
	for i := 0; i < len(str); i++ {
		_, err = strconv.Atoi(str[i : i+1])
		if err == nil {
			temp = temp + str[i:i+1]
		} else {
			num[j], _ = strconv.Atoi(temp)
			glyph[j] = str[i : i+1]
			j++
			temp = ""
		}
	}
	num[j], _ = strconv.Atoi(temp)
	for i := 0; i < j; i++ { // Нужно реализовать чтобы он делал опереции поочередно (сначала умножение потом остально и тд)
		switch glyph[i] {
		case "/":
			num[i], num[i+1] = div(num[i], num[i+1]), 0 //Н
		case "*":
			num[i], num[i+1] = mul(num[i], num[i+1]), 0
		}
	}
	for i := 0; i < j; i++ {
		switch glyph[i] {
		case "+":
			result = result + sum(num[i], num[i+1])
		case "-":
			result = result + sub(num[i], num[i+1])
		}
	}
	fmt.Println(result)
}
