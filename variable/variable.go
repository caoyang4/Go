package main

import "fmt"

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

}
