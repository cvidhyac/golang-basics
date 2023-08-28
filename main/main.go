package main

import "fmt"

func main() {
	fmt.Println("Print a string")

	var value = "Printing from a variable reference \n"
	fmt.Print(value)

	var firstname = "FirstName"
	var lastname = "LastName"
	fmt.Print("Welcome concatenated", firstname, " ", lastname)

	var grade = 42
	fmt.Printf("Grades formatted as decimal: %d \n", grade)

	bitwiseOperator()
	shiftOperator()
}

/**
* Example for bitwise AND and OR operator
 */
func bitwiseOperator() {
	var x, y = 100, 90
	fmt.Println(x & y)
	fmt.Println(x | y)
}

func shiftOperator() {
	var x, y = 100, 90
	fmt.Println((x + y) >> 2)
}
