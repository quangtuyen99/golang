package main

import "fmt"

func main() {

	var dimension int

	// Taking input from user
	fmt.Scanln(&dimension)

	for i := dimension; i > 0; i-- {
		count := i

		for count > 0 {
			fmt.Print("*")
			count -= 1
		}
		fmt.Println()
	}
}
