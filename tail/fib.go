package tail

func Fibonacci(num int) int {
	return num
}

func fib(num int, next int, result int) int {
	if num == 0 {
		return result
	}
	return fib(num-1, result, result+next)
}
