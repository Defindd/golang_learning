package main

import "fmt"

const (
	one = 1 + 2*iota
	three
	five
	seven
	nine
	eleven
)

func main() {
	a := 51
	switch {
	case a > 0:
		if a%2 == 0 {
			break
		}
		fmt.Println("Odd positive value received")
		fallthrough
	case a < 0:
		fmt.Println("Negative value received")

	default:
		fmt.Println("Default value handling")
	}
}
