package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%p \n", a)
	b := a[1:2]
	b = append(b, 15)
	fmt.Printf("%p\n", b)
	b = append(b, 75)
	fmt.Println(a, b)
}
