package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Getenv("USER")
	course := "Docker Deep Dive"

	fmt.Println("\n Hi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\n You are currently watching", course)

}
func changeCourse(course *string) string{
	*course = "go fundamentals"
	fmt.Println("Trying to change course", *course)
	return *course
}