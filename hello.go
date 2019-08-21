package main

import (
	"fmt"
)

const k string = "key"

func main() {
	fmt.Println("Hello go!")

	var i = 1
	fmt.Print(i)
	i = 3
	j := 2
	fmt.Print(j)

	var x = "hell"

	fmt.Print(x)
	x = "hello"

	fmt.Print("\n")

	fmt.Println("7.00/1.00", 7.00/1.0)
	fmt.Println(true)

	a, b, c := 1, 2, 3

	fmt.Println(a + b + c)

	fmt.Println(len(x) + a)

	var e int
	fmt.Println(e)
	var x1 string
	fmt.Println(x1)

	fmt.Println(k)

	const n = 50000

	fmt.Println(int64(n))
	fmt.Println(n / 2e30)
}
