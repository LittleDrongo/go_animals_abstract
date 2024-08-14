package animals

type Cat struct {
	AnimalImpl
}

func (c *Cat) Meow() {
	println("Мяу!!")
}
