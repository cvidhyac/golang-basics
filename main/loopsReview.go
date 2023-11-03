package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Examples of the 5 loop patterns to understand the for construct")

	fmt.Println("1. As a three component loop")
	threeComponentLoop()

	fmt.Println("2. As a while loop")
	asAWhileLoop()

	fmt.Println("3a. As a foreach range loop, iterate on list")
	forEachRangeLoopList()

	fmt.Println("3b. As a for each Range loop, iterate on map")
	forEachRangeLoopMap()

	fmt.Println("4. As an infinite loop, the example is commented out for brevity")
	//comment infinite loop as it runs forever
	//asInfititeLoop()

	fmt.Println("5a. Exit loop with continue")
	exitLoopWithContinueKeyword()

	fmt.Println("5b. Exit loop with break keyword")
	exitLoopWithBreakKeyword()
}

func exitLoopWithBreakKeyword() {
	countryMap := make(map[string]string)
	countryMap = map[string]string{
		"India":     "Delhi",
		"Canada":    "Ottawa",
		"USA":       "WashingtonDC",
		"Australia": "Sydney",
	}

	var country string
	var capital string
	for key, value := range countryMap {
		if strings.EqualFold("Ottawa", value) {
			country = key
			capital = value
			break
		}
	}
	fmt.Printf("Found country : %v, Capital %v", country, capital)
}

func threeComponentLoop() {

	maxTimes := 10
	for i := 0; i < maxTimes; i++ {
		fmt.Printf("I will execute %v times \n", maxTimes)
	}
}

func asAWhileLoop() {

	maxTimes := 10
	for maxTimes <= 0 {
		fmt.Printf("While I am not zero, i execute, current run : %v", maxTimes)
		maxTimes = maxTimes - 1
	}
}

func forEachRangeLoopList() {

	alphaList := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Print a list .. ")
	for index := range alphaList {
		fmt.Println(alphaList[index])
	}
}

func forEachRangeLoopMap() {
	maps := map[string]int{
		"Key1": 1,
		"Key2": 2,
		"Key3": 3,
	}
	for key, value := range maps {
		fmt.Printf("Key: %v , Value: %v \n", key, value)
	}
}

func exitLoopWithContinueKeyword() {
	values := []string{"A", "B", "C", "D", "E"}
	var findMe string
	for index := range values {
		if strings.EqualFold(values[index], "D") {
			continue
		}
		findMe = values[index]
	}
	if len(findMe) == 0 {
		fmt.Println("Value Not found")
	} else {
		fmt.Printf("Found %v \n", findMe)
	}

}
func asInfiniteLoop() {

	status := true
	for status {
		fmt.Println("Run till the end of this world")
	}
}
