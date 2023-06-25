package main

import (
	"fmt"
)

// 闭包函数
func Adder() func(int) int {
	n := 0
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	f := Adder()
	fmt.Println("f(1)=", f(1))
	fmt.Println("f(2)=", f(2))
	fmt.Println("f(3)=", f(3))
	fmt.Println("f(4)=", f(4))
}
