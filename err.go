package main

import (
	"errors"
	"fmt"
)

func Increment(i int) (int, error) {

	if i < 10 {
		return 0, errors.New("negative number")
	}

	return i + 1, nil
}

func main() {
	num := 5

	if inc, err := Increment(num); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inc)
	}
}
