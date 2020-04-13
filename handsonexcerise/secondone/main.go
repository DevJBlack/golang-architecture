package main

import "fmt"

type person struct {
	first string
}

type dbService struct {
	a accessor
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func (ds dbService) get(n int) (person, error) {
	p := ds.a.retrieve(n)
	if p.first == "" {
		return person{}, fmt.Errorf("No Person with n of %d", n)
	}
	return p, nil
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

type mysql map[int]person
type postgres map[int]person

func (m mysql) save(n int, p person) {
	m[n] = p
}

func (m mysql) retrieve(n int) person {
	return m[n]
}

func (pg postgres) save(n int, p person) {
	pg[n] = p
}

func (pg postgres) retrieve(n int) person {
	return pg[n]
}

func main() {

	dbsql := mysql{}
	dbpg := postgres{}

	p1 := person{
		first: "Devin",
	}

	p2 := person{
		first: "Timp",
	}

	put(dbsql, 1, p1)
	fmt.Println("this is going to Sql DB", get(dbsql, 1))
	put(dbsql, 2, p2)
	fmt.Println("This is from SQL DB", get(dbsql, 2))

	put(dbpg, 1, p1)
	fmt.Println("This is going to Postgres DB", get(dbpg, 1))
	put(dbpg, 2, p2)
	fmt.Println("This is from Postgres DB", get(dbpg, 2))

}
