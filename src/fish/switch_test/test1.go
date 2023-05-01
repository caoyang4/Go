package main

import (
	"fmt"
)

func getByte(char byte) byte {
	return char + 1
}

func main() {

	var alpha byte
	fmt.Scanf("%c", &alpha)

	switch alpha {
	case 'a':
		fmt.Println("alpha a")
	case 'b':
		fmt.Println("alpha b")
	case 'c':
		fmt.Println("alpha c")
	case 'd', 'e', 'f':
		fmt.Println("alpha d|e|f")
	default:
		fmt.Println("others")
	}

	fmt.Println("===============")

	switch getByte(alpha) - 1 {
	case 'a':
		fmt.Println("alpha a")
	case 'b':
		fmt.Println("alpha b")
	case 'c':
		fmt.Println("alpha c")
		// 穿透一层
		fallthrough
	case 'd', 'e', 'f':
		fmt.Println("alpha d|e|f")
	default:
		fmt.Println("others")

	}

	fmt.Println("=============")

	switch {
	case alpha == 'a':
		fmt.Println("alpha a")
	case alpha >= 'x':
		fmt.Println("alpha after x")
	case alpha > 'b':
		fmt.Println("alpha after b and before x")

	}

}
