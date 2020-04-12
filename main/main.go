package main

import "fmt"

// Concret is person
type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent -- this is my name", sa.first)
}

// any Type that has the methods specified by an interface
// is also of the interface type
// an interface says
// if you have these methods, then you're my type.
type human interface {
	speak()
}

func main() {
	// Abstract type is p1
	p1 := person{
		first: "Devin",
	}
	sal := secretAgent{
		person: person{
			first: "DevJBlack",
		},
		ltk: true,
	}

	fmt.Printf("%T\n", p1)
	// in go a VALUE can be of more than one TYPE
	// in this example, p1 is both TYPE person and human
	var x, y human
	x = p1
	y = sal
	x.speak()
	y.speak()
}
