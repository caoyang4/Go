package utils

import (
	"fmt"
)

// 首字母小写表示 private
var name = "this is a utils pkg"

// 首字母大写表示 public
var Human = "mao"

func Test() {
	fmt.Println(name)
}
