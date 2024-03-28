package main

import "fmt"

func main() {

	var dimension int

	// Taking input from user
	fmt.Scanln(&dimension)

	starLastRow := dimension*2 - 1

	// Star will print in range (padLeft, padRight)
	padLeft := starLastRow / 2
	padRight := starLastRow / 2

	padReverseLeft := 1
	padReverseRight := starLastRow

	// Triangle
	for i := 0; i < dimension; i++ {
		for j := 0; j < starLastRow; j++ {
			if j >= padLeft && j <= padRight {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		padLeft -= 1
		padRight += 1

		fmt.Println()
	}

	// Reverse triangle
	for i := 0; i < dimension; i++ {
		for j := starLastRow; j >= 0; j-- {
			if j >= padReverseLeft && j <= padReverseRight {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		padReverseLeft += 1
		padReverseRight -= 1

		fmt.Println()
	}
}
