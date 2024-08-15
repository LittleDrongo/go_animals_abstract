package animals

type Cat struct {
	AnimalAbstract
}

func (c *Cat) Meow() {
	println("Мяу!!")
}
