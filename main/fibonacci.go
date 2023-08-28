package main

import "fmt"

func findFibonacci(num int) int {
	if num < 2 {
		return num
	}
	return findFibonacci(num-1) + findFibonacci(num-2)
}

func main() {
	var fibArray []int
	for i := 0; i <= 10; i++ {
		fibonacciNum := findFibonacci(i)
		fibArray = append(fibArray, fibonacciNum)
	}
	fmt.Println(fibArray)
}
