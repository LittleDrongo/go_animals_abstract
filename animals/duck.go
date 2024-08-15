package animals

import (
	"fmt"
)

type Duck struct {
	AnimalAbstract
}

func (d *Duck) Quack() {
	fmt.Println("Кря!!!")
}

func (d *Duck) Eat() {
	fmt.Printf("Утка %v кушает по другому \n", d.name)
}
