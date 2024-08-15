package animals

import "fmt"

type Duck struct {
	AnimalAbstract
}

func (d *Duck) Quack() {
	fmt.Println("Кря!!!")
}

/*
// Собственная реализация Eat, будет отличаться от Animal!
func (d *Duck) Eat() {
	fmt.Println("ЭТО ЖИВОТНОЕ КУШАЕТ ПО ДРУГОМУ, ПОТОМУ ЧТО ЭТО УТКА С СОБСТВЕННОЙ ФУНКЦИЕЙ EAT")

}
*/
