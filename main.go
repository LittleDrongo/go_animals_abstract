package main

import (
	"app/animals"
	"fmt"
)

func main() {

	cat := &animals.Cat{}
	dog := &animals.Dog{}
	duck := &animals.Duck{}

	// cat.SetName("Солнышко")
	cat.New("Солнышко")
	cat.Tail = 50

	dog.New("Пятнышко")
	duck.New("Скрудж")

	// Обобщеные функции
	cat.Eat()
	dog.Eat()

	duck.Eat() //@override функции //TODO

	// Индивидуальные методы.
	cat.Meow()
	dog.Bark()

	fmt.Println(cat)

}

/* //README
Принципы ООП

1. Инкапсуляция
2. Полиморфизм
3. Наследование
4. Абстракция
5. Расширение*

*/
