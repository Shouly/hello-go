package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name  string
	age   int
	phone int64
}

type age int

type student struct {
	person
	score int
}

func Older(p1, p2 person) (person, int) {
	return p1, p1.age - p2.age
}

func main() {

	var x age

	fmt.Println(x)

	var p person

	p.name = "Sh"

	p.age = 100

	fmt.Println(p)

	n := person{"h", 1, 2}
	fmt.Println(n)

	var s student

	fmt.Println(s)

	var f int = 4
	reflect.ValueOf(f).Elem().Field(0)
}
