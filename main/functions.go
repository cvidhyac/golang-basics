package main

import "fmt"

func main() {

	var (
		myFunc = func(s string) {
			fmt.Println("My Value", s)
		}
	)

	myFunc("Hello")

}
