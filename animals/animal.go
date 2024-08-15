package animals

import "fmt"

// Определяем интерфейс Animal
type Animal interface {
	Eat()
	SetName(string)
}

type AnimalAbstract struct {
	name string
}

func (a *AnimalAbstract) SetName(name string) {
	a.name = name
}

/*
func (a *AnimalImpl) GetName() string {
	return a.name
}
*/

func (a *AnimalAbstract) Eat() {
	fmt.Printf("Животное %v кушает \n", a.name)
}
