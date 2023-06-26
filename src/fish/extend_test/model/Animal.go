package model

type Animal struct {
	color  string
	weight float64
	name   string
}

func (animal *Animal) GetColor() string {
	return animal.color
}

func (animal *Animal) SetColor(color string) {
	animal.color = color
}

func (animal *Animal) GetWeight() float64 {
	return animal.weight
}

func (animal *Animal) SetWeight(weight float64) {
	animal.weight = weight
}

func (animal *Animal) GetName() string {
	return animal.name
}

func (animal *Animal) SetName(name string) {
	animal.name = name
}
