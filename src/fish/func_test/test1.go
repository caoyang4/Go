package func_test

import "fmt"

func Cal(n1 float64, n2 float64, oper byte) float64 {
	var result float64
	switch oper {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		result = n1 / n2
	default:
		return 0.0
	}
	return result
}

func RecursiveTest(n int) {
	if n > 2 {
		n--
		RecursiveTest(n)
	}
	fmt.Println("n=", n)
}
