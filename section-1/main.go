package main

import (
	"LEARN-GO/section-1/test"
	"fmt"
	"strings"

	a "github.com/iancoleman/strcase"
)

func main() {
	fmt.Println("Main Go")

	var greeting string
	greeting = "This variabel"
	fmt.Println(greeting)

	greeting2 := "This short"
	fmt.Println(greeting2)

	greeting3 := strings.ToLower("LOWER VARIABLE")
	fmt.Println(greeting3)

	greeting4 := a.ToSnake("SNAKE VARIABLE")
	fmt.Println(greeting4)

	fmt.Println(test.FirstFunction("world"))
	text, lengg := test.SecondFunction("world")
	fmt.Println(text)
	fmt.Println(lengg)

	name := "John Doe"
	address := "Jakarta"
	fmt.Printf("name : %s\naddress : %s\n", name, address)
}
