package mongo

import "https://github.com/DevJBlack/golang-architecture/tree/master/datastorage/storage/mongo"

type Db map[int]architecture.Person

func (m Db) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) architecture.Person {
	return m[n]
}
