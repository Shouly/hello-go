package main

import "fmt"

func main() {

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
