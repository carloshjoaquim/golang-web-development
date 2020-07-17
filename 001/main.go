package main

import "fmt"

type human interface {
	speak()
}

type agent interface {
	shoot()
}

type person struct {
	firstName string
	lastName string
	age int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, `says, "Hello ! I'm a person.!"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.lastName, sa.firstName, `says, "Hello..."`)
}

func (sa secretAgent) shoot() {
	fmt.Println("Agent shooting !")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person {
		"Maria",
		"Rock",
		28,
	}

	sa1 := secretAgent{
		person{"James",
			"Bond",
			42,
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)
}
