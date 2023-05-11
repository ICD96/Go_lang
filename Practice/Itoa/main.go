package main

import (
	"fmt"
)

func itoa(n int) string {
	s := ""
	t := 0
	for n != 0 {
		t = n % 10
		n = n / 10
		switch t {
		case 0:
			s = "0" + s
		case 1:
			s = "1" + s
		case 2:
			s = "2" + s
		case 3:
			s = "3" + s
		case 4:
			s = "4" + s
		case 5:
			s = "5" + s
		case 6:
			s = "6" + s
		case 7:
			s = "7" + s
		case 8:
			s = "8" + s
		case 9:
			s = "9" + s
		}
	}
	return s
}

func main() {
	c := itoa(14)
	fmt.Println(c)
}
