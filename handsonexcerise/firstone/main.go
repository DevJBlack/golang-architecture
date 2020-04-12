package main

import "fmt"

type area interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	p1 := square{
		sideLength: 10,
	}

	p2 := triangle{
		base:   10,
		height: 10,
	}

	printArea(p1)
	printArea(p2)
}

func printArea(a area) {
	fmt.Println(a.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}
