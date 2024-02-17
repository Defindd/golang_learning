package main

import (
	"fmt"
)

// var name string = "Nikita"
type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Married(husband Person) {
	p.LastName = husband.LastName
}

type WeekDay int32

func main() {
	const (
		mon WeekDay = iota + 1
		tue
		wed
		thu
		fri
		sat
		sun
	)
	fmt.Println(mon, tue, thu, sun)
}
