package main

import "fmt"

type AnimalInterface interface {
	Say()
}

type Dog struct { // Doesn't implement AnimalInterface but say about it
	AnimalInterface
	name string
}

func AnimalSay(animal AnimalInterface) {
	animal.Say()
}

func main() {
	dog := Dog{name: "Snow"}
	fmt.Printf("Dog: %+v\n", dog) // let's see what we have in dog
	dog.Say()                     // <- And compilator let you use method from AnimalInterface
}
