package main

import "fmt"

func main() {

	var dimension int

	// Taking input from user
	fmt.Scanln(&dimension)

	for i := 1; i <= dimension; i++ {
		count := 1

		for count <= i {
			fmt.Print(count)
			count += 1
		}
		fmt.Println()
	}
}
