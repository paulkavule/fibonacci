package head

import "errors"

func Fibonacci(num int) (int, error) {
	if num <= 0 {
		return 0, errors.New("input has to be greater than zero")
	}
	fib := fib(num)
	return fib, nil
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
