package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in f", r)
		}
	}()

	fmt.Println("calling g")
	g(0)
	fmt.Println("return from g")
}

func g(i int) {

	if i > 3 {
		fmt.Println("Panicking!")
		panic(i)
	}

	defer fmt.Println("Defer in g", i)

	fmt.Println("Printing in g", i)

	g(i + 1)

}

func main() {
	f()
	fmt.Println("return from f")
}
