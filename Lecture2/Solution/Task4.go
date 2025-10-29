package main

import "fmt"

func main() {
	// Check if a number is even or odd
	var number int
	fmt.Println("Enter an integer:")
	fmt.Scanln(&number)
	
	if number % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Grading system
	var score int
	fmt.Println("Enter the score (0-100):")
	fmt.Scanln(&score)
	switch {
	case score >= 80 && score <= 100:
		fmt.Println("A+")
	case score >= 75 && score <= 79:
		fmt.Println("A")
	case score >= 70 && score <= 74:
		fmt.Println("A-")
	case score >= 65 && score <= 69:
		fmt.Println("B+")
	case score >= 60 && score <= 64:
		fmt.Println("B")
	case score >= 55 && score <= 59:
		fmt.Println("B-")
	case score >= 50 && score <= 54:
		fmt.Println("C+")
	case score >= 45 && score <= 49:
		fmt.Println("C")
	case score >= 40 && score <= 44:
		fmt.Println("D")
	case score >= 0 && score <= 39:
		fmt.Println("F")
	default:
		fmt.Println("Invalid score")
	}

}