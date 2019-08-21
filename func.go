package main

import "fmt"

func main() {

	a, b := 1, 2
	fmt.Println(maxAndMin(a, b))
	fmt.Println(a, b)
	fmt.Println(math(a, b))

	n(1, 3, 4, 5, 6)

	idx(&a)

	fmt.Println(a)
}

func maxAndMin(a, b int) (x, y int) {

	a = a + 1

	if a > b {
		return a, b
	}

	return b, a
}

func math(a, b int) (sum, diff, multi, division int) {
	return a + b, a - b, a * b, a / b
}

func n(args ...int) {
	for i := range args {
		fmt.Println(args[i])
	}
}

func idx(i *int) {
	*i = *i + 1
	fmt.Println(i)
}

func add(n int) int {
	if n == 0 {
		return 0
	}
	return n + add(n-1)
}
