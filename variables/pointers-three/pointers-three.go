package main

import (
	"fmt"
)

func main() {

	name := "Nikhil"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\n you are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "First look: Native Docker Clustering"
	fmt.Println("\n Trying to change your course to", *course)
	return *course
}