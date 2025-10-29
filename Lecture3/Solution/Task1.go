package main

import (
    "fmt"
    "math"
)
func main() {
	fmt.Printf("Is 5 prime?  %t\n", isPrime(5))  // Output: true
	fmt.Printf("Is 4 prime?  %t\n", isPrime(4))  // Output: false
	fmt.Printf("Is 7 prime?  %t\n", isPrime(7))  // Output: true
	fmt.Printf("Is 1 prime?  %t\n", isPrime(1))  // Output: false
	fmt.Printf("Is -10 prime? %t\n", isPrime(-10)) // Output: false
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }else if n ==2 {
        return true
    }
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
	    if n % i == 0 {
	        return false
	    }
	}
    return true
}