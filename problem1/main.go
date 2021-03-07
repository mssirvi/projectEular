package main

import (
	"fmt"
)

func solution() int {
	sum := 0
	for i := 3; i < 1_000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
func main() {
	fmt.Println(solution())
}
