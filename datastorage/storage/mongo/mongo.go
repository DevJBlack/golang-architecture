package mongo

import "C:\Users\Zendragon\DevMountain\DevMountain-school\udemy-classes\udemy-go-course\exploringGo\golang-architecture\datastorage\storage\mongo"

type Db map[int]architecture.Person

func (m Db) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) architecture.Person {
	return m[n]
}
