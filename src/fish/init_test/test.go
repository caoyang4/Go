package main

import (
	"fmt"
)

// 全局变量先初始化
var age = ofAge()

func ofAge() int {
	fmt.Println("ofAge()")
	return 18
}

// init可以完成初始化工作
func init() {
	fmt.Println("init()")
}

func main() {

	fmt.Println("main()")
	fmt.Println("age=", age)

}
