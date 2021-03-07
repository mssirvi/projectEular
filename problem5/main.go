package main

import "fmt"

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(x int, y int) int {
	return (x * y) / gcd(x, y)
}
func main() {
	res := 1
	for i := 1; i <= 20; i++ {
		res = lcm(res, i)

	}
	fmt.Println(res)
}
