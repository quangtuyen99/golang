package main

import "fmt"

func main() {

	var dimension int
	var line string

	// Taking input from user
	fmt.Scanln(&dimension)

	for i := 0; i < dimension; i++ {
		line += "*"
	}

	for i := 0; i < dimension; i++ {
		fmt.Println(line)
	}
}
