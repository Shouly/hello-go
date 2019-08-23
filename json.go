package main

import (
	"encoding/json"
	"fmt"
)

type r struct {
	pageIndex int      `json:"page"`
	fruits    []string `json:"fruits"`
}

func main() {

	//encode
	m := map[string]int{"apple": 5, "x": 1}

	fmt.Println(m)

	j, _ := json.Marshal(m)

	fmt.Println(j)

	//decode

	str := `{"page":1, "fruits":["a","b"]}`

	res := new(r)

	err := json.Unmarshal([]byte(str), &res)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.fruits)
}
