package main

import "fmt"

func sum(a []int, c chan int) {

	total := 0

	for _, v := range a {
		total += v
	}

	c <- total
}

func fibonacci(n int, c chan int) {

	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func fib(c, quit chan int) {

	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y)

	////
	fmt.Println("///")

	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("///")

	ch, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fib(ch, quit)
}
