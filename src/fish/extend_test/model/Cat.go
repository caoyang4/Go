package model

import "fmt"

// 继承
type Cat struct {
	Animal
	kind string
}

func (cat *Cat) GetKind() string {
	return cat.kind
}

func (cat *Cat) SetKind(kind string) {
	cat.kind = kind
}

func (cat *Cat) String() string {
	desc := fmt.Sprintf("name: %v, weight %v, color: %v, kind: %v", cat.name, cat.weight, cat.color, cat.kind)
	return desc
}
