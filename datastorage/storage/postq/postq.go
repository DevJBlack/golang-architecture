package postq

import "https://github.com/DevJBlack/golang-architecture/tree/master/datastorage/storage/postq"

type Db map[int]architecture.Person

func (pq Db) Save(n int, p architecture.Person) {
	pq[n] = p
}

func (pq Db) Retrieve(n int) architecture.Person {
	return pq[n]
}
