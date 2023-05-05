package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	// 当程序执行到 defer 时，暂时不执行，会将 defer 之后的语句压入到独立的栈中
	// 当函数执行结束，再从 defer 栈，按照先入后出的方式出栈执行
	// 当压入栈，也会把响应的值拷贝进去
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	n1++
	n2++
	n := n1 + n2
	fmt.Println("n=", n)
	return n
}
func main() {

	r := sum(1, 2)
	fmt.Println("r=", r)

}
