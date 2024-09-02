package main

import "fmt"

// Объявляем Human
type Human struct {
	name string
	age  int
}

func (h Human) Speak() {
	fmt.Printf("%s is %d years old\n", h.name, h.age)
}

type Action struct {
	Human // Встраивание
}

func main() {
	a := Action{
		Human{
			name: "Baga",
			age:  20,
		},
	}
	a.Speak() // вызываем метод Speak, который принадлежит Human
}
