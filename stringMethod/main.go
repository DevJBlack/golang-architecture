package main

import "fmt"

type person struct {
	first string
}

func (p person) String() string {
	return fmt.Sprintf("My name is %s", p.first)
}

func main() {
	p1 := person{
		first: "Devin",
	}
	fmt.Println(p1)
}
