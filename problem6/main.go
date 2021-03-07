package main

import (
	"fmt"
)

func squareSum(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func sumSquares(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)

}
func main() {
	fmt.Println(sumSquares(100) - squareSum(100))
}
