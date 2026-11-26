package main

import "fmt"

func fib(n int) {
	if n <= 1 {
		fmt.Print(n, " ")
		return
	}
	var a, b int = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		fmt.Print(b, " ")
	}
}

func main() {
	fmt.Println("Fibonacci sequence:")
	fib(10)
}
