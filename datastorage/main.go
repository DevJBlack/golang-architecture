package main

import "fmt"

type person struct {
	first string
}

type mongo map[int]person
type postg map[int]person

func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrive(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) retrive(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrive(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrive(n)
}

func main() {
	dbm := mongo{}
	dbp := postg{}

	p1 := person{
		first: "Devin",
	}

	p2 := person{
		first: "James",
	}

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println("this is saving to Mongo DB", get(dbm, 1))
	fmt.Println("This is retriving from Mongo DB", get(dbm, 2))

	put(dbp, 1, p1)
	put(dbp, 2, p2)
	fmt.Println("This is saving to PostQ DB", get(dbp, 1))
	fmt.Println("This is retriving from PostQ BD", get(dbp, 2))
}
