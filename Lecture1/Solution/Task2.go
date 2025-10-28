package main

import "fmt"

func main() {	
	var name string
	var grade_1, grade_2, grade_3 float64
	fmt.Print("Enter student's name: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter Grades separated by space: ")
	fmt.Scan(&grade_1, &grade_2, &grade_3)

	average := (grade_1 + grade_2 + grade_3) / 3
	
	fmt.Printf("Student Name: %s\n", name)
	fmt.Printf("Grades: %.2f, %.2f, %.2f\n", grade_1, grade_2, grade_3)
	fmt.Printf("Average Grade: %.2f\n", average)

}
