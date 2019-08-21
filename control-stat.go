package main

import (
	"fmt"
	"math"
)

func main() {

	i := 1
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i < 100; i++ {
		fmt.Println(int8(i))
	}

	fmt.Println(math.Abs(-100))

	for {
		break
	}

	for i := 0; i < 100; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	var (
		pi int8
		x1 string
	)

	fmt.Println(pi)
	fmt.Println(x1)

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)
}
