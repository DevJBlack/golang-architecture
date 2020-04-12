package architecture

import "fmt"

// Person is how the architecture package stores a person
type Person struct {
	first string
}

// Accessor is how to store or retrieve a person.
// When retriving a person, if they do not exist, return the zero value
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

// PersonService struct
type PersonService struct {
	a Accessor
}

// NewPersonService function
func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

// Get function retrieves info from Db
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.first == "" {
		return Person{}, fmt.Errorf("No Person with n of %d", n)
	}
	return p, nil
}

// Put function saves a person in the Db
func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

// Get function returns person from Db
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
