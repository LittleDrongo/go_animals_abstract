package main

import "app/animals"

func main() {

	cat := &animals.Cat{}
	dog := &animals.Dog{}
	duck := &animals.Duck{}

	duck.SetName("Скрудж")
	cat.SetName("Солнышко")
	dog.SetName("Пятнышко")

	// Обобщеные функции
	cat.Eat()
	dog.Eat()
	duck.Eat() //@override функции

	// Индивидуальные методы.
	cat.Meow()
	dog.Bark()
	duck.Quack()
}
