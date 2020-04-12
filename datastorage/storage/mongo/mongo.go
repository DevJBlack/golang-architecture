package mongo

import "github.com/DevJBlack/golang-architecture"

type Db map[int]architecture.Person

func (m Db) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) architecture.Person {
	return m[n]
}
