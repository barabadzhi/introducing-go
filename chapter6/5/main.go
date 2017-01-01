package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(0)) // 0
	fmt.Println(fib(1)) // 1
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(9)) // 34
}
