package main

import (
	"fmt"
)

func main() {

	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("年龄大于18")
	} else {
		fmt.Println("小屁孩")
	}

	// golang 支持在 if 中直接定义变量
	if i := age; i > 100 {
		fmt.Println(i)
	}

}
