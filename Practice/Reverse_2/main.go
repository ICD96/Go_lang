package main

import "fmt"

func main() {
	str := "Привет"
	runes := []rune(str)
	result := make([]rune, 0, len(runes))
	for i := len(runes); i != 0; i-- {
		fin = append(result, runes[i-1])
	}
	fmt.Println(string(fin))
}
