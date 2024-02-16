package main

import (
	"fmt"
	"reflect"
)

// var name string = "Nikita"
type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Married(husband Person) {
	p.LastName = husband.LastName
}
func main() {
	k := 15
	defer func() {
		fmt.Printf("Первый дефер %d\n", k)
		k += 25
		defer func() {
			fmt.Printf("второй дефер %d\n", k)
		}()
	}()
	k = 55
	k++
	/* fmt.Printf("My name is %s\n", name)

	fmt.Println("Hello world!")
	fmt.Print("Hello world2!")
	fmt.Println("Hello world3!")
	//var hello string = "hello world 4"
	b1 := true
	//var hello2 string = "qwerty"
	//fmt.Printf("Type: %T Value:%v\n", hello, hello)
	fmt.Printf("%+d, %o %t, ", -5, 175, b1)
	fmt.Print("123")
	fmt.Println("c\u00e1t")
	var a rune = 128512
	println(string(a))
	frm, wlc := SayHi("Nikita", "Prilukov")
	fmt.Printf("%s,%s", frm, wlc)*/

	GetAr := func(x, y int16) int16 {
		return x + y
	}
	IsMultiple := GetMultipleFunction(15, 2, 3, GetAr)
	fmt.Println(reflect.TypeOf(IsMultiple))
	if IsMultiple() {
		fmt.Println("Кратно!")
	} else {
		fmt.Println("Не кратно!")
	}

	i := 10
	for {
		i++
		fmt.Println(i)
		if i > 15 {
			break
		}
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ты лох")
		}
	}()

}

/*
	func SayHi(firstname, secondname string) (formatedString string, welcome string) {
		formatedString = fmt.Sprintf("Hi %q", firstname+" "+secondname)
		welcome = "welcome to c1dy's ass"
		return
	}
*/
func GetMultipleFunction(a, b, divider int16, GetArythm func(x, y int16) int16) (DividerFunc func() bool) {
	DividerFunc = func() bool {
		a++
		fmt.Println(a)
		return GetArythm(a, b)%divider == 0
	}
	return
}
func GetAge(age int) int {
	return age + 5
}
func AddNum(num *int) {
	*num++
}
