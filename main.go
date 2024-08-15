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

/* //README
Принципы ООП

1. Инкапсуляция
2. Полиморфизм
3. Наследование
4. Абстракция
5. Расширение*

*/
