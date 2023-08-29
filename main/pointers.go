package main

import "fmt"

func addressAndDereferenceOperator() {
	i := 10
	fmt.Printf("Address of the datatype and accessing pointer %T %v \n", &i, &i)
	fmt.Printf("Dereferencing the pointer, %T, %v", *(&i), *(&i))
}

func main() {
	addressAndDereferenceOperator()
}
