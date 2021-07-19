package main

import "fmt"

func main() {
	fmt.Println("==Fibonacci==")
	for i := 0; i < 15; i++ {
		fmt.Print(" ", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
