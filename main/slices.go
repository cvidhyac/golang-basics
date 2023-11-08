package main

import (
	"fmt"
	"reflect"
	"slices"
)

/*
 * 1. Understand high and low bounds
 * 2. Iteration starts from zero to nth element
 * 3. When a limit is not specified, slicing uses '0' for lower boundary, or max size of the array for upper boundary.
 */
func main() {
	initSlice()
	printSlice()
	isArrayEqual := compareSlice([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
	fmt.Printf("Result of array comparision %v \n", isArrayEqual)
	sortArrOrSlice()

}

func initSlice() {
	parentArr := []int{1, 2, 3, 4, 5, 6}

	aStringArr := make([]string, 2, 4) //arrayType, len, capacity

	parentSlice := parentArr[2:4]
	fmt.Printf("ParentSlice %v \n", parentSlice)

	aStringArr = []string{"a", "b", "c", "d"}
	fmt.Printf("A string array sliced upto 2 positions : %v, length : %v, capacity : %v \n",
		aStringArr, len(aStringArr), cap(aStringArr))
}
func printSlice() {

	fmt.Println("Start printing slices .. ")

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

/*
*
* 3 ways to compare slices
* 1. compare length, if length is different return false
* 2. Run a for loop, and compare elements, if not matching return false
* 3. use reflect.DeepEqual in tests to recursively compare a slice. This is not recommended for production code.
 */
func compareSlice(a, b []int) bool {

	//approach 1
	if len(a) != len(b) {
		return false
	}

	//approach 2
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	//approach 3
	if reflect.DeepEqual(a, b) {
		return true
	} else {
		return false
	}

}

func sortArrOrSlice() {
	sortMe := []int{21, 51, 22, 54, 33}
	slices.Sort(sortMe)
	fmt.Printf("Sorted Array %v \n", sortMe)
}
