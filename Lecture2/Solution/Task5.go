package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Enter coefficients a, b and c: ")
	fmt.Scanln(&a, &b, &c)
	
	if a == 0 && b != 0 {
		x := -c / b
		fmt.Printf("Linear equation root: %.2f\n", x)
		return
	}else if b == 0 {
		if c == 0 {
			fmt.Println("Infinite solutions")
		} else {
			fmt.Println("No real solutions")
		}
	} else {
		discriminant := math.Pow(b, 2) - 4*a*c
		if discriminant > 0 {
			root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
			root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
			fmt.Printf("Two real and distinct roots: %.2f and %.2f\n", root1, root2)
		} else if discriminant == 0 {
			root := -b / (2 * a)
			fmt.Printf("One real root: %.2f\n", root)
		} else {
			fmt.Println("No real solutions")
		}
	}
}