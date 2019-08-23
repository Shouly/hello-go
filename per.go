package main

import (
	"fmt"
	person2 "hello-go/person"
)

func main() {
	p := person2.Person{Name: "s"}
	fmt.Println(person2.Description(p))
}
