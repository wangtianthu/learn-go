package main

import (
	"fmt"
)

type Person interface {
	Eat() string
	Speak() string
}

type Man struct {
	name string
	food string
}

func (man Man) Eat() string {
	return fmt.Sprintf("%s is eating %s", man.name, man.food)
}

func (man Man) Speak() string {
	return fmt.Sprintf("%s said: I am eating %s", man.name, man.food)
}

func main() {
	man := Man{"Jack", "sandwich"}
	var p Person
	p = &man
	fmt.Printf("-- %s\n", p.Eat())
	fmt.Printf("-- %s\n", p.Speak())
}
