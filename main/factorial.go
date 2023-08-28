package main

import "fmt"

func findFactorial(num int) int {
	if num == 1 {
		return num
	}
	return num * findFactorial(num-1)
}

func main() {
	fmt.Println(findFactorial(5))
}
