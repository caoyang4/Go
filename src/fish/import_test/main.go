package main

import (
	"fish/func_test"
	"fish/import_test/utils"
	"fmt"
)

var (
	// 全局匿名函数
	MyVar = func(n1 int, n2 int) int { return n1 * n2 }
)

func mod(p *int) {
	*p += 1
	fmt.Println("p=", *p)
}

// 函数也可作为形参类型
func myFunc(funcVar func(float64, float64, byte) float64, n1 float64, n2 float64, oper byte) float64 {
	return funcVar(n1, n2, oper)
}

type funcType func(float64, float64, byte) float64

func myFunc1(funcVar funcType, n1 float64, n2 float64, oper byte) float64 {
	return funcVar(n1, n2, oper)
}

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

	fmt.Println("Cal")
	result := func_test.Cal(2.1, 3.4, '*')
	fmt.Println(result)

	fmt.Println("myFunc")
	result = myFunc(func_test.Cal, 3.1, 4.2, '+')
	fmt.Println(result)

	fmt.Println("myFunc1")
	result = myFunc1(func_test.Cal, 3.1, 4.2, '/')
	fmt.Println(result)

	fmt.Println("RecursiveTest")
	func_test.RecursiveTest(5)

	fmt.Println("Fib")
	fmt.Println("Fib(10)=", func_test.Fib(10))

	fmt.Println("Series")
	fmt.Println("Series(10)=", func_test.Series(10))

	n := 1
	mod(&n)
	// 函数也是一类数据类型
	x := mod
	fmt.Println("n=", n)
	x(&n)
	fmt.Println("n=", n)

	fmt.Println("Sum")
	r := func_test.Sum(1, 2, 3, 4, 5)
	fmt.Println(r)

	fmt.Println("Swap")
	n1 := 2
	n2 := 3
	fmt.Printf("n1=%v, n2=%v \n", n1, n2)
	func_test.Swap(&n1, &n2)
	fmt.Printf("n1=%v, n2=%v \n", n1, n2)

	// 匿名函数
	fmt.Println("匿名函数")
	sum := func(n1 int, n2 int) int {
		return n1 * n2
	}(2, 5)
	fmt.Println("sum=", sum)
	fmt.Println("multiple=", MyVar(2, 3))

}
