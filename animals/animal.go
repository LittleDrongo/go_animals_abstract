package animals

import "fmt"

// Определяем интерфейс Animal
type Animal interface {
	Eat()
	SetName(string)
}

type AnimalImpl struct {
	name string
}

func (a *AnimalImpl) SetName(name string) {
	a.name = name
}

/*
func (a *AnimalImpl) GetName() string {
	return a.name
}
*/

func (a *AnimalImpl) Eat() {
	fmt.Printf("Животное %v кушает \n", a.name)
}
