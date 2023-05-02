package func_test

func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func Series(n int) int {
	if n == 1 {
		return 3
	}
	return 2*Series(n-1) + 1
}

func Sum(n int, args ...int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func Swap(n1, n2 *int) {
	*n1 = (*n1) ^ (*n2)
	*n2 = (*n1) ^ (*n2)
	*n1 = (*n1) ^ (*n2)
}
