package main

import (
	"fmt"

	"github.com/tiket/section2/internal/subpackage2"
)

func main() {
	fmt.Println("File main")
	subpackage2.SubPackage2Function()
	mainTwo()
	mainThree()
	mainFour()
}

func tambah(angka *int) {
	*angka = *angka + 1
}

func mainTwo() {
	fmt.Println("main two")
	angka := 13
	fmt.Println("Sebelum: ", angka)
	tambah(&angka)
	fmt.Println("Sesudah: ", angka)

	subpackage2.SubPackage2Function()
}

func mainThree() {
	var slice []int
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	for index, item := range slice {
		fmt.Println(index, item)
	}

	slice = append(slice, 17)
	fmt.Println(slice)

	str := "Word from String"
	fmt.Println(str[1:3])

	m := map[string]string{
		"new": "value",
	}
	//replace init map
	m = make(map[string]string)
	m["key"] = "test"
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}
}

type Prayer interface {
	Pray()
}

func (p *Person) Pray() {
	fmt.Println("Zzzzzzzz")
}

type Live struct {
	Blessing string
}

type Person struct {
	Live

	Name string
	Age  int
}

func mainFour() {
	person := Person{
		Name: "John Doe",
		Age:  22,
		Live: Live{
			Blessing: "bless bless bless",
		},
	}

	fmt.Printf("%+v\n", person)

	person.Birthday()

	fmt.Printf("%+v\n", person)

	fmt.Println(person.Live.Blessing)
	person.Pray()
}
