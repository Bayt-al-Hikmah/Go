package main
import "fmt"

func main() {
	// Factorial calculation
	var n int
	fmt.Print("Enter a positive integer to calculate its factorial: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
	} else {
		factorial := 1
		i := 1
		for i <= n {
			factorial *= i
			i++
		}
		fmt.Printf("Factorial of %d is %d\n", n, factorial)
	}
	
}