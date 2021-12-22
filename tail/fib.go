package tail

import "errors"

func Fibonacci(num int) (int, error) {
	if num <= 0 {
		return 0, errors.New("number should be greater than 0")
	}
	fib := fib(num, 0, 1)
	return fib, nil
}

func fib(num int, next int, result int) int {
	if num == 0 {
		return result
	}
	return fib(num-1, result, result+next)
}
