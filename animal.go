package main

import "fmt"

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func main() {
	var a animal
	a = snake{Type: "x", Poisonous: true}
	fmt.Println(a.description())
}
