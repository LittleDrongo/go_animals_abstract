package animals

// Структура Dog, которая встраивает AnimalImpl
type Dog struct {
	AnimalImpl
}

// Метод Bark для Dog
func (d *Dog) Bark() {
	println("Гав!!")
}
