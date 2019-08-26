package main

import "fmt"

type X struct {
	name string
}

func main() {

	////

	x1 := X{name: "1"}

	x1.name = "x1"

	fmt.Println(x1)

	var x2 X

	x2.name = "x2"

	fmt.Println(x2)

	x3 := &X{}

	x3.name = "x3"

	fmt.Println(*x3)

	x4 := new(X)

	x4.name = "x4"

	fmt.Println(*x4)

	/////

	p := new(int)

	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	o := 0

	fmt.Println(o)
	fmt.Println(&o)

	var p2 *int
	i := 0
	p2 = &i

	fmt.Println(p2)
	fmt.Println(*p2)

	pf(p2)
	pf(*p2)

}

func pf(f interface{}) {
	fmt.Println(f)
}
