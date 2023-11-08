package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4}
	arr = pushElements(arr, 5)
	arr = pushElements(arr, 6)
	fmt.Println(arr)       // {0, 1, 2, 3, 4, 5, 6}
	arr = popElements(arr) // {0, 1, 2, 3, 4, 5}
	fmt.Println(arr)
}

func popElements(arr []int) []int {
	lastIndex := len(arr) - 1
	return arr[:lastIndex]
}
func pushElements(arr []int, element int) []int {
	mySlice := append(arr, element)
	return mySlice
}
