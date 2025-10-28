package main

import "fmt"

func main() {	

	var radius float64
	const Pi = 3.14159

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanln(&radius)
	fmt.Printf("The surface area of the circle is: %.2f\n", Pi*radius*radius)
}