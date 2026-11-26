package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}

func main() {
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scan(&num)
	if isPrime(num) {
		fmt.Println("The number is prime.")
	} else {
		fmt.Println("The number is not prime.")
	}
}
