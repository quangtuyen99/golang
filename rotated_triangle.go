package main

import "fmt"

func main() {

	var dimension int

	// Taking input from user
	fmt.Scanln(&dimension)

	for i := 0; i < dimension; i++ {
		count := 0

		for count <= i {
			fmt.Print("*")
			count += 1
		}
		fmt.Println()
	}

	// Reverse
	for i := 0; i < dimension-1; i++ {
		countReverse := dimension - 1

		for countReverse > i {
			fmt.Print("*")
			countReverse -= 1
		}
		fmt.Println()
	}
}
