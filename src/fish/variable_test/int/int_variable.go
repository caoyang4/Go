package main

import "fmt"
import "unsafe"

var g1 = 6
var g2 = "global"

var (
	g3 = 2
	g4 = "young"
)

func main() {

	fmt.Println("g1", g1, "g2", g2)
	fmt.Println("g3", g3, "g4", g4)

	var i int = 1
	fmt.Println("i=", i)

	var j int
	fmt.Println("j=", j)

	var pi = 3.14
	fmt.Println("pi=", pi)

	name := "james"
	fmt.Println("name=", name)

	var n1, n2, n3 int
	fmt.Println(n1, " ", n2, " ", n3)

	var m1, m2, m3 = 100, 5.12, "kobe"
	fmt.Println(m1, " ", m2, " ", m3)

	k1, k2, k3 := 1, 1.23, "ok"
	fmt.Println(k1, " ", k2, " ", k3)

	x := 1.2
	x = 1
	fmt.Println("x=", x)

	y := 1
	z := 1
	fmt.Println("y+z=", y+z)

	// -128 - 127
	var ii1 int8 = 127
	fmt.Println(ii1)

	// 0 - 255
	var ii2 uint8 = 255
	fmt.Println(ii2)

	x1 := int32(ii2)
	fmt.Printf("type of x1 is %T \n", x1)

	var x2 float32 = float32(ii2)
	fmt.Printf("type of x2 is %T \n", x2)

	x3 := 128
	fmt.Printf("type of x3 is %T, size: %d \n", x3, unsafe.Sizeof(x3))

	var x4 int
	fmt.Println("x4=", x4)
}
