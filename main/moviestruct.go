package main

import "fmt"

type Movie struct {
	name    string
	ratings float32
}

func getMovie(name string, ratings float32) (m Movie) {
	m = Movie{
		name:    name,
		ratings: ratings}
	return
}

func main() {

	mov := getMovie("xyz", 2.0)
	fmt.Println(mov.name)
	fmt.Println(mov.ratings)

	mov1 := getMovie("xyz", 2.1)
	mov2 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov1)
	movies = append(movies, mov2)

	for _, value := range movies {
		fmt.Println(value)
	}
}
