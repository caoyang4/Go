package main

import (
	"fmt"
)

func main() {
	fmt.Println("loop test")
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("============")
	var j = 1
	for j <= 10 {
		fmt.Println("j=", j)
		j++
	}

	for {
		fmt.Println("月读")
		if j++; j > 15 {
			break
		}
	}
	fmt.Println("============")
	var name = "young"
	for i := 0; i < len(name); i++ {
		fmt.Printf("index=%d, val=%c \n", i, name[i])
	}

	fmt.Println("============")
	// for-range 方式
	for index, val := range "sjtu" {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

	var slogan = "闷声发大财"
	slgs := []rune(slogan)
	for i := 0; i < len(slgs); i++ {
		fmt.Printf("index=%d, val=%c \n", i, slgs[i])
	}

}
