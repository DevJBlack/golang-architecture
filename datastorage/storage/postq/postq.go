package postq

import "C:\Users\Zendragon\DevMountain\DevMountain-school\udemy-classes\udemy-go-course\exploringGo\golang-architecture\datastorage\storage\postq"

type Db map[int]architecture.Person

func (pq Db) Save(n int, p architecture.Person) {
	pq[n] = p
}

func (pq Db) Retrieve(n int) architecture.Person {
	return pq[n]
}
