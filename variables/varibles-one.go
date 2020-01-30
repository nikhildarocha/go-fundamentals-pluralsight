// package main


// var name string = "Nigel"
// var course string = "Docker Deep Dive"
// var module float64 = 2.6

package main

import (
	"fmt"
	"reflect"
)

var (
	name string // Name of Subscriber
	course string // Name of current course
	module float64 // current place in course

)

func main () {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
}