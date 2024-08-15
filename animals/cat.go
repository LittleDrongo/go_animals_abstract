package animals

type Cat struct {
	AnimalAbstract
	Tail float64
}

func (c *Cat) Meow() {
	println("Мяу!!")
}
