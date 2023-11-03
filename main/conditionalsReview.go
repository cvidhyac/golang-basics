package main

import (
	"fmt"
	"math/rand"
	t "time"
)

func main() {
	if lessonLearned := true; lessonLearned {
		fmt.Println("Great job! You can continue on to the next exercise.")
	} else {
		fmt.Println("Practice makes perfect.")
	}
	// Create more conditions below!
	for i := 0; i < 10; i++ {
		fmt.Println(createGrade(printRandom()))
	}

}

func printRandom() int {
	rand.NewSource(t.Now().UnixNano())
	return rand.Intn(99)
}

func createGrade(num int) string {

	var grade string
	if num > 80 {
		grade = "A"
	} else if num > 70 && num <= 80 {
		grade = "B"
	} else if num > 60 && num <= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	return grade
}
