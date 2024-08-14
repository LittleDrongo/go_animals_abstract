package main

import "fmt"

func main() {

	cat := Cat{AnimalImpl{Name: "Ketty"}}
	dog := Dog{AnimalImpl{Name: "Doggy"}}

	cat.Eat()
	dog.Eat()

	cat.Meow()
	dog.Bark()
}

// Структура Cat, которая встраивает AnimalImpl
type Cat struct {
	AnimalImpl
}

// Метод Meow для Cat
func (c *Cat) Meow() {
	fmt.Println("Meow!")
}

// Структура Dog, которая встраивает AnimalImpl
type Dog struct {
	AnimalImpl
}

// Метод Bark для Dog
func (d *Dog) Bark() {
	fmt.Println("Bark!")
}

// Определяем интерфейс Animal с методом Eat()
type Animal interface {
	Eat()
}

// Базовая структура AnimalImpl, которая реализует интерфейс Animal
type AnimalImpl struct {
	Name string
}

// Реализация метода Eat для AnimalImpl
func (a *AnimalImpl) Eat() {
	fmt.Printf("The animal %v is eating\n", a.Name)
}
