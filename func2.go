package main

import "fmt"

func main() {

	arr := []int{1, 3, 4, 5, 5}

	fmt.Println(filter(arr, isOdd))
	fmt.Println(filter(arr, isOk))
}

type check func(int) bool

func isEven(v int) bool {
	return v%2 == 0
}

func isOdd(v int) bool {
	return !isEven(v)
}

func isOk(v int) bool {
	return false
}

func filter(arr []int, f check) []int {
	var result []int

	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}
