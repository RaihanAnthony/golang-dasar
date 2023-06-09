package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(HasName HasName){
	fmt.Println("hello", HasName.getName())
}

type person struct {
	Name string
}

func (person person) getName() string {
	return person.Name
}

type animal struct {
	name string
}

func (animal animal) getName() string {
	return animal.name
}

func main() {
	person := person{Name: "raihan"}
	SayHello(person)

	animal := animal{
		name: "kucing",
	}
	SayHello(animal)
	
}