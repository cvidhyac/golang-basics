package main

import "fmt"

func calcArea(radius float64) (float64, string) {
	return 3.14 * radius * radius, "area"
}

func calcPerimeter(radius float64) (float64, string) {
	return 2 * 3.14 * radius, "perimeter"
}

func calcDiameter(radius float64) (float64, string) {
	return 2 * radius, "diameter"
}

func mapFunction(query int) func(radius float64) (float64, string) {
	queryToFunction := map[int]func(radius float64) (float64, string){
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return queryToFunction[query]
}

func findValue(radius float64, execFunction func(radius float64) (float64, string)) (float64, string) {
	return execFunction(radius)
}
func main() {
	myCalculation, calculationType := findValue(10, mapFunction(1))
	fmt.Printf("The result of my calculation is %v and the calculationType is %v \n :", myCalculation, calculationType)
	fmt.Println("Thank you!!")
}
