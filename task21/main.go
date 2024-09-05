package main

import "fmt"

func main() {
	cat := NewCatAdapter(&Cat{})
	cat.Reaction()
	dog := NewDogAdapter(&Dog{})
	dog.Reaction()
}

type AnimalAdapter interface {
	Reaction()
}

type Cat struct{}

func (c *Cat) Say() {
	fmt.Println("I am a cat")
}

type CatAdapter struct {
	*Cat
}

func (c *CatAdapter) Reaction() {
	c.Say()
}

func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

type Dog struct{}

func (d *Dog) Say() {
	fmt.Println("I am a dog")
}

type DogAdapter struct {
	*Dog
}

func (d *DogAdapter) Reaction() {
	d.Say()
}

func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}
