package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	fmt.Println("goroutine test")
	// 协程
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
