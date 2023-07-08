package main

import (
	"fmt"
	utils "practice_maps/Practice_Maps/Utils"
)

func main() {
	mape := utils.Construct()
	fmt.Println(mape)
	utils.Set(mape, "1", "One")
	fmt.Println(mape)
}
