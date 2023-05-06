package main

import (
	"fmt"
)

// go中数组属于值类型，默认值传递，因此进行值拷贝，数组间不会互相影响

func test1(arr [3]int) {
	arr[0] = 6
	fmt.Println("test1 arr6=", arr)
}

func test2(arr *[3]int) {
	(*arr)[1] = 6
	fmt.Println("test2 arr7=", *arr)
}

func main() {
	fmt.Println("array test")
	var arr [3]int
	arr[0] = 1
	fmt.Println("arr=", arr)

	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1=", arr1)
	var arr2 = [3]string{"a", "b", "c"}
	fmt.Println("arr2=", arr2)

	var arr3 = [...]string{"a", "b", "c", "d"}
	fmt.Println("arr3=", arr3)

	// 指定下标
	var arr4 = [...]string{3: "a", 2: "b", 1: "c", 0: "d"}
	fmt.Println("arr4=", arr4)

	arr5 := [...]string{3: "a", 2: "b", 1: "c", 0: "d"}
	fmt.Println("arr5=", arr5)

	// 遍历数组

	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3: index: %d, value: %v \n", i, arr3[i])
	}

	for i, v := range arr5 {
		fmt.Printf("arr5: index: %d, value: %v \n", i, v)
	}

	arr6 := [3]int{1, 2, 3}
	fmt.Println("before main arr6=", arr6)
	test1(arr6)
	fmt.Println("after main arr6=", arr6)

	arr7 := [3]int{1, 2, 3}
	fmt.Println("before main arr7=", arr7)
	test2(&arr7)
	fmt.Println("after main arr7=", arr7)

}
