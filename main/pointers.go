package main

import "fmt"

func addressAndDereferenceOperator() {
	i := 10
	fmt.Printf("Address of the datatype and accessing pointer %T %v \n", &i, &i)
	fmt.Printf("Dereferencing the pointer, %T, %v \n", *(&i), *(&i))
}

func initPointers() {
	value := "hello world"
	secondValue := &value
	fmt.Println(*secondValue)
	manipulatePointer(&value)
	fmt.Println(*secondValue)
}

func manipulatePointer(value *string) {
	*value = "new world"
}
func main() {
	addressAndDereferenceOperator()
	initPointers()
}
