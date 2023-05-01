package main

import (
	"fish/func_test"
	"fish/import_test/utils"
	"fmt"
)

func main() {
	fmt.Println("xx")
	utils.Test()
	fmt.Println(utils.Human)

	a := 1
	b := 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a)
	fmt.Println(b)

	result := func_test.Cal(2.1, 3.4, '*')
	fmt.Println(result)
}
