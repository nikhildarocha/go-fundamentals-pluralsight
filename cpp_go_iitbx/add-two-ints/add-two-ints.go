package main

import (
	"fmt"
)

func addTwoInts(a, b int) int {
	return a + b
}
func main() {
	var A, B, C int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&A, &B)
	C = addTwoInts(A, B)
	fmt.Println("The sum is: ", C)

}
