package main

import (
	"fmt"
)

func main() {
	fmt.Println("map test")

	// 声明
	var aMap map[string]string = make(map[string]string, 10)
	// 在使用前，需要用 make 分配空间
	aMap["james"] = "23"
	fmt.Println("aMap=", aMap)

	var bMap = map[int]string{1: "x", 2: "y"}
	fmt.Println("bMap=", bMap)

	stu := make(map[string]map[string]string)
	stu["james"] = make(map[string]string)
	stu["james"]["name"] = "lebron"
	stu["james"]["num"] = "23"
	stu["kobe"] = make(map[string]string)
	stu["kobe"]["name"] = "brant"
	stu["kobe"]["num"] = "24"
	fmt.Println("stu=", stu)

	msg, exist := stu["paul"]
	if exist {
		fmt.Println(msg)
	} else {
		fmt.Println("not exist")
	}

}
