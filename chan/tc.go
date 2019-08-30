package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	defer func() {
		close(c)
		fmt.Println("chan c closed.")
	}()

	c <- 1

	runtime.Gosched()

	fmt.Println(<-c)
}
