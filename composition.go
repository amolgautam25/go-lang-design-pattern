package main

import "fmt"

// decalre the Animal Struct
type Animal struct {
	id string
}

func (a *Animal) walk() {
	fmt.Println("walking -->", a.id)
}

// create a dog struct
type Dog struct {
	id string
	Animal
}

func (d *Dog) talk() {
	fmt.Println("I bark", d.id)
}

// create a cat struct
type Cat struct {
	id string
	Animal
}

func (c *Cat) talk() {
	fmt.Println("I meow", c.id)
}

func main() {

	fmt.Println("Hello World")

	cat1 := Cat{id: "cat1", Animal: Animal{id: "cat1"}}
	dog1 := Dog{id: "dog1", Animal: Animal{id: "dog1"}}

	fmt.Print("Now, lets make them talk \n")
	cat1.talk()
	dog1.talk()

	fmt.Println("Now lets make them walk")
	cat1.walk()
	dog1.walk()

}
