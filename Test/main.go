package main

import "fmt"

func main() {
	str := "1234"
	num, err := atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

func atoi(str string) (int, error) {
	var num int
	for _, v := range str {
		digit := int(v - '0')
		if digit < 0 || digit > 9 {
			return 0, fmt.Errorf("invalid digit: %v", v)
		}
		num = num*10 + digit
	}
	return num, nil
}
