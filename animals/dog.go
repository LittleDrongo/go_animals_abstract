package animals

// Структура Dog, которая встраивает AnimalImpl
type Dog struct {
	AnimalAbstract
}

// Метод Bark для Dog
func (d *Dog) Bark() {
	println("Гав!!")
}
