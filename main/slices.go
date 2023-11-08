package main

import "fmt"

/*
 * 1. Understand high and low bounds
 * 2. Iteration starts from zero to nth element
 * 3. When a limit is not specified, slicing uses '0' for lower boundary, or max size of the array for upper boundary.
 */
func main() {
	printSlice()
}
func printSlice() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s)
	s1 := s[:] // {0, 1, 2, 3, 4}
	fmt.Println(s1)
	s2 := s[:2] // {0, 1}
	fmt.Println(s2)
	s3 := s[1:2] // {1}
	fmt.Println(s3)
	s4 := s[:3] // {0, 1, 2}
	fmt.Println(s4)
	s5 := s[3:] // {3, 4}
	fmt.Println(s5)

}
