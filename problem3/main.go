package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(x int) int {
	for i := int(math.Sqrt(float64(x))); i >= 2; i-- {
		if x%i == 0 && isPrime(i) {
			return i
		}

	}
	return 1
}
func isPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
