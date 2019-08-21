package main

import "fmt"

type Rectangle struct {
	width, height float32
}
type Circle struct {
	radius float32
}

func (r Rectangle) area() float32 {

	return r.height * r.width
}

func (c Circle) area() float32 {
	return c.radius
}

func main() {

	r := Rectangle{1.3, 2.89}
	fmt.Println(r.area())
	c := Circle{1.2}
	fmt.Println(c.area())
}
