package main

import (
	"fmt"
)

func main() {
	fmt.Println("begin")
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Println("name=", name)

	fmt.Scanln(&age)
	fmt.Println("age=", age)

}
