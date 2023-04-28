package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i = 100
	fmt.Printf("type of i is: %T \n", i)

	var j int64 = 100
	fmt.Printf("type of j is: %T, size of is %d \n", j, unsafe.Sizeof(j))

	var k float32 = 3.14
	fmt.Println(k)

	var m float64 = 2.33
	fmt.Println(m)

	// 默认 float64 双精度
	var n = 3.14
	fmt.Printf("type of n is %T \n", n)

	x := .123
	fmt.Printf("type of x is %T \n", x)

	y := 'a'
	fmt.Printf("type of y is %T \n", y)
	fmt.Println(y)

	var z byte = 'b'
	fmt.Printf("type of z is %T \n", z)
	fmt.Println(z)
	fmt.Printf("z = %c \n", z)

	var zi int = '好'
	fmt.Printf("zi = %c, zi = %d \n", zi, zi)

	zj := 120
	fmt.Printf("zj = %c \n", zj)

	var z1 = 10 + 'a'
	fmt.Println(z1)
}
