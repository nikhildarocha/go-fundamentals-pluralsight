package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Nikhil"
	course := "GO Fundamentals"
	module := 3.2
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Memory address of *module* variable is", &module)
	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)
}