package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	// todo 递归和全局变量 来创建树结构
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(4))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(4))
}
