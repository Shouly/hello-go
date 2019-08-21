package main

import "fmt"

type person struct {
	name  string
	age   int
	phone int64
}

type student struct {
	person
	score int
}

func Older(p1, p2 person) (person, int) {
	return p1, p1.age - p2.age
}

func main() {
	var p person

	p.name = "Sh"

	p.age = 100

	fmt.Println(p)

	n := person{"h", 1, 2}
	fmt.Println(n)

	var s student

	fmt.Println(s)
}
