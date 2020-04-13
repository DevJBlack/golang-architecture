package main

import "fmt"

type user struct {
	first string
}

type mongo map[int]user

func (m mongo) save(n int, u user) {
	m[n] = u
}

func (m mongo) retriever(n int) user {
	return m[n]
}

type harddrive map[int]user

// Same syntax as the accessor interface, we just pass in a recever of either harddrive or mongo
func (hd harddrive) save(n int, u user) {
	hd[n] = u
}

// Same syntax as the accessor interface, we just pass in a recever of either harddrive or mongo
func (hd harddrive) retriever(n int) user {
	return hd[n]
}

// the save and retriever are the exactly the same as the functions
type accessor interface {
	save(n int, u user)
	retriever(n int) user
}

func put(a accessor, n int, u user) {
	a.save(n, u)
}

func get(a accessor, n int) user {
	return a.retriever(n)
}

func main() {
	storage := harddrive{}
	db := mongo{}

	u1 := user{
		first: "Devin",
	}

	u2 := user{
		first: "Jesse",
	}

	put(storage, 1, u1)
	put(storage, 2, u2)

	fmt.Println("Get from storage", get(storage, 1))
	fmt.Println("Get from storage", get(storage, 2))

	put(db, 1, u1)
	put(db, 2, u2)

	fmt.Println("Get from db", get(db, 1))
	fmt.Println("Get from db", get(db, 2))
}
