package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Action struct {
	Human
}

func (a Action) Talk() {
	fmt.Println("Привет!Меня зовут", a.name)
}

func main() {
	first_person := Action{
		Human: Human{
			name:  "Юра",
			age:   23,
			phone: "777",
		},
	}

	first_person.Talk()
	fmt.Println(first_person)
}
