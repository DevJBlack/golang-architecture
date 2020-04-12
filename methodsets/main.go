package main

import "fmt"

type person struct {
	first string
}

type secertAgent struct {
	person
	ltk bool
}

func (p *person) speak() {
	fmt.Println("From a person", p.first)
}

func (sa *secertAgent) speak() {
	fmt.Println("this is not my real name", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Jimmy",
	}

	sal := secertAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	saySomething(&p1)
	saySomething(&sal)

}
