package main

import (
	"fmt"
)

func main() {

	i := 1
	fmt.Println("addr of i is ", &i)

	var ptr *int = &i
	fmt.Println(ptr, " ", *ptr)

	var ptr1 = &ptr
	fmt.Printf("i = %d \n", **ptr1)

	_j := 1
	fmt.Println("_j = ", _j)
}
