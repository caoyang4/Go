package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	// go中字符串不可变
	name := "hello"
	fmt.Println(name)

	// 反引号输出字符串
	code := `func main() {
	var i = 100
	fmt.Printf("type of i is: %T \n", i)

	var j int64 = 100
	fmt.Printf("type of j is: %T, size of is %d \n", j, unsafe.Sizeof(j))
	`
	fmt.Println(code)

	name += " world!"
	fmt.Println(name)

	num := 6
	str := fmt.Sprintf("int=%d", num)
	fmt.Println(str)
	str = fmt.Sprintf("float=%f", 1.23)
	fmt.Println(str)
	isOKay := true
	str = fmt.Sprintf("bool=%t", isOKay)
	fmt.Println(str)
	char := 'c'
	str = fmt.Sprintf("char=%c", char)
	fmt.Println(str)
	str1 := "okay"
	str = fmt.Sprintf("string=%q", str1)
	fmt.Println(str)

	isOKay, _ = strconv.ParseBool("true")
	fmt.Println(isOKay)

	numStr := "123456"
	num1, _ := strconv.ParseInt(numStr, 10, 0)
	fmt.Printf("num1=%d type is %T, size = %d\n", num1, num1, unsafe.Sizeof(num1))

	floatNum := "123.456"
	f1, _ := strconv.ParseFloat(floatNum, 64)
	fmt.Printf("f1=%f type is %T, size = %d\n", f1, f1, unsafe.Sizeof(f1))

}
