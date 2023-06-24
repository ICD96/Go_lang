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
	str := "1+2*3+4" //Работает не со всеми условиями нужно будет попробовать 3*2+1*2 или 1+2*3+5
	var err error
	num := []int{}
	glyph := []string{}
	temp := ""
	temp2 := 0
	check := false
	j := 0
	for i := 0; i < len(str); i++ {
		_, err = strconv.Atoi(str[i : i+1])
		if err != nil {
			if string(str[i]) == "*" || string(str[i]) == "/" {
				glyph = append(glyph, string(str[i]))
				temp2, _ = strconv.Atoi(temp)
				num = append(num, temp2)
				fmt.Println(num[j])
				j++
				check = true
			} else {
				check = false
			}
			temp = ""
		} else {
			temp = temp + string(str[i])
		}
	}
	if check {
		temp2, _ = strconv.Atoi(temp)
		num = append(num, temp2)
		fmt.Println(num[j])
		j++
	}
	check = false
	temp = ""
	for i := 0; i < len(str); i++ {
		_, err = strconv.Atoi(str[i : i+1])
		if err != nil {
			if string(str[i]) == "+" || string(str[i]) == "-" {
				glyph = append(glyph, string(str[i]))
				temp2, _ = strconv.Atoi(temp)
				num = append(num, temp2)
				check = true
				fmt.Println(num[j])
				j++
			} else {
				check = false
			}
			temp = ""
		} else {
			temp = temp + string(str[i])
		}
	}
	if check {
		temp2, _ = strconv.Atoi(temp)
		num = append(num, temp2)
		fmt.Println(num[j])
		j++
	}
	for i := 0; i < len(glyph); i++ {
		switch glyph[i] {
		case "/":
			num[i+1] = div(num[i], num[i+1])
		case "*":
			num[i+1] = mul(num[i], num[i+1])
		case "+":
			num[i+1] = sum(num[i], num[i+1])
		case "-":
			num[i+1] = sub(num[i], num[i+1])
		}
	}
	fmt.Println(num[len(num)-1])
}
