package person

import "fmt"

type Person struct {
	Name string
}

func Description(p Person) string {
	return fmt.Sprint("name is ", p.Name)
}

func secretName() string {
	return "can't"
}
