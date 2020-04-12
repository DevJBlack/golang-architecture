package main

import (
	"fmt"
	"github.com/DevJBlack/golang-architecture"
	"github.com/DevJBlack/golang-architecture/datastorage/mongo"
	"github.com/DevJBlack/golang-architecture/datastorage/postq"

)

func main() {
	dbm := mongo.Db{}
	dbp := postq.Db{}

	p1 := architecture.Person{
		First: "Devin",
	}

	p2 := architecture.Person{
		First: "James",
	}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println("this is saving to Mongo DB", architecture.Get(dbm, 1))
	fmt.Println("This is retriving from Mongo DB", architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println("This is saving to PostQ DB", architecture.Get(dbp, 1))
	fmt.Println("This is retriving from PostQ BD",architecture.Gget(dbp, 2))
}
