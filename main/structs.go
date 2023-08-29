package main

import "fmt"

type Circle struct {
	radius int
	x      int
	y      int
}

func initStructWithNew() *Circle {
	circle := new(Circle)
	circle.radius = 10
	circle.x = 5
	circle.y = 10
	return circle
}

func initStructInline() Circle {
	return Circle{15, 16, 17}
}
func main() {

	newCircle := initStructWithNew()
	fmt.Printf("%+v \n", *newCircle)

	inlineCircle := initStructInline()
	fmt.Printf("%+v \n", inlineCircle)

}
