package main

import (
	"fmt"
)

func main() {
	num := new(int)
	*num = 1

	fmt.Printf("num=%v, *num=%v, &num=%v", num, *num, &num)
}
