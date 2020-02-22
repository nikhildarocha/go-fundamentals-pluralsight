package main

import "fmt"

func main() {

	//firstRank := "39"
	//secondRank := "614"

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than" + " than second course")
	} else if firstRank > secondRank {
		fmt.Println("Oh dear, your first course is absymall!")
	} else {
		fmt.Println("Both courses are either the same or something weird is going on")
	}
}
