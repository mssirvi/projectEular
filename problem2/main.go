package main

import (
	"fmt"
)

const fibNumber = 4_000_000 // number literal prefix for readability

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fib(n-1) + fib(n-2)
}
func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
func main() {
	sum := 0
	for i := 1; fib(i) < fibNumber; i++ {
		num := fib(i)
		if isEven(num) {
			sum += num
		}
	}
	fmt.Println(sum)
}
