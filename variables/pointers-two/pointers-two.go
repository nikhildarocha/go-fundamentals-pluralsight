package main

import (
	"fmt"
)

func main() {
	name := "Nikhil"
	course := "Go fundamentals"
	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(course)
	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course string) string {
	course = "First look: Native Docker Clustering"
	fmt.Println("Trying to change your course to", course)
	return course
}