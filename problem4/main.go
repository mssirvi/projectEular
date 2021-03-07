package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	real := x
	reverse := 0
	for x != 0 {
		reverse = (reverse * 10) + x%10
		x = x / 10
	}
	return reverse == real
}
func main() {
	const max = 999
	maxProduct := 1
	for i := max; i > 99; i-- {
		for j := i; j > 99; j-- {
			product := i * j
			if isPalindrome(product) && product > maxProduct {
				maxProduct = product
			}
		}
	}

	fmt.Println(maxProduct)
}
