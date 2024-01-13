package main

//var name string = "Nikita"

func main() {
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
	IsMultiple := GetMultipleFunction(10, 2, 3, GetAr)
	if IsMultiple() {
		println("Кратно!")
	} else {
		println("Не кратно!")
	}
}

/*func SayHi(firstname, secondname string) (formatedString string, welcome string) {
	formatedString = fmt.Sprintf("Hi %q", firstname+" "+secondname)
	welcome = "welcome to c1dy's ass"
	return
}*/
func GetMultipleFunction(a, b, divider int16, GetArythm func(x, y int16) int16) (DividerFunc func() bool) {
	DividerFunc = func() bool {
		return GetArythm(a, b)%divider == 0
	}
	return
}
