package animals

import (
	"fmt"
	"time"
)

type Animal interface {
	Eat()
	SetName(string)
}

type AnimalAbstract struct {
	name    string
	addDate time.Time
}

func (a *AnimalAbstract) GetName() string {
	return a.name
}

func (a *AnimalAbstract) Eat() {
	fmt.Printf("Животное %v кушает \n", a.name)
}

func (a *AnimalAbstract) New(name string) {
	a.name = name
	a.addDate = time.Now()

}

func (a AnimalAbstract) String() string {
	return "Животное: " + a.name + "Добавлено:" + a.addDate.GoString()

}
