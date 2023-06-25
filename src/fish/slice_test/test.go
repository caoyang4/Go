package main

import (
	"fmt"
)

// 切片是引用类型
func main() {

	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr=", arr)
	sliceArr := arr[1:3]
	fmt.Println("sliceArr=", sliceArr)
	fmt.Printf("addr of arr[1]: %v, addr of sliceArr[0]: %v \n", &arr[1], &sliceArr[0])
	sliceArr[0] = 6

	fmt.Println("after arr=", arr)
	fmt.Println("after sliceArr=", sliceArr)

	var slice = make([]string, 5, 10)
	slice[0] = "x"
	fmt.Println("slice=", slice)

	slice1 := []float64{1.2, 2.3, 3.4}
	fmt.Printf("slice1=%v type:%T\n", slice1, slice1)

	var strSlice = []string{"kobe", "james", "jordan"}
	strSlice = append(strSlice, "wade")
	strSlice = append(strSlice, []string{"curry", "nash"}...)
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	for i, v := range arr {
		fmt.Printf("index=%v, value=%v \n", i, v)
	}

	// 切片拷贝
	var sliceA = []int{1, 2, 3}
	var sliceB = make([]int, 5, 10)
	copy(sliceB, sliceA)
	fmt.Println("sliceA=", sliceA)
	fmt.Println("sliceB=", sliceB)
	sliceA[0] = 6
	sliceB[0] = 7
	fmt.Println("sliceA=", sliceA)
	fmt.Println("sliceB=", sliceB)

	str := "Ilove中国"

	// 这样遍历中文字符会乱码
	for i := 0; i < len(str); i++ {
		fmt.Printf("index=%v, value=%c \n", i, str[i])
	}
	fmt.Println()
	for i, v := range str {
		fmt.Printf("index=%v, value=%c \n", i, v)
	}

	fmt.Println()
	slice3 := []rune(str)
	for i, v := range slice3 {
		fmt.Printf("index=%v, value=%c \n", i, v)
	}

	var slice4 []int = []int{2, 1}
	slice4[0] = 1
	fmt.Println("slice4=", slice4)
}
