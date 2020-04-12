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

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

type personService struct {
	a accessor
}

func (ps personService) get(n int) (person, error) {
	p := ps.a.retrieve(n)
	if p.first == "" {
		return person{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
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

	ps := personService{
		a: dbm,
	}

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println("this is saving to Mongo DB", get(dbm, 1))
	fmt.Println("This is retriving from Mongo DB", get(dbm, 2))

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3))

	put(dbp, 1, p1)
	put(dbp, 2, p2)
	fmt.Println("This is saving to PostQ DB", get(dbp, 1))
	fmt.Println("This is retriving from PostQ BD", get(dbp, 2))
}
