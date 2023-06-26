package main

import (
	"fish/extend_test/model"
	"fmt"
)

func main() {

	cat := model.Cat{}
	cat.SetColor("yellow")
	cat.SetWeight(6.18)
	cat.SetName("lulu")
	cat.SetKind("金渐层")

	fmt.Println(cat)

}
