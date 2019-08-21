package main

import "fmt"

func main() {

	// array
	var a [10]int

	a[0] = 1
	a[1] = 2

	a[9] = 3
	fmt.Println(a)

	b := [2]int{1, 2}

	fmt.Println(b)

	//slice

	fmt.Println("/////")

	slice := []int{1, 2, 3, 4, 5}

	var arr = slice[0:1]

	fmt.Println(slice)
	fmt.Println(arr)

	arr = append(arr, 2, 3, 4, 5, 6)
	fmt.Println(arr)

	fmt.Println("/////")

	s := make([]string, 2)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"

	fmt.Println(len(s))
	s = append(s, "c")
	s = append(s, "d")

	fmt.Println(s)

	var x []string
	copy(x, s)

	fmt.Println(x)
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// map
	fmt.Println("/////")

	var m map[string]int

	m = make(map[string]int)

	m["x"] = 1

	fmt.Println(m)
	fmt.Println(m["y"])
	fmt.Println(len(m))

	n := map[string]int{"x": 1, "y": 2}

	fmt.Println(n)

	mn, ok := n["mn"]

	if !ok {
		fmt.Println(mn, ok)
	}

	for k, v := range n {
		fmt.Println(k, v)
	}
}
