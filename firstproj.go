package main

import (
	"fmt"
)

func main() {
	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	var qwerty = &weekTemp
	qwerty[2] = 123
	fmt.Println(weekTemp, *qwerty)
}
