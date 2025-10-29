
package main

import "fmt"


func main() {
	fmt.Printf("Binary of 5 is %d\n", toBinary(5))   // Output: 101
	fmt.Printf("Binary of 10 is %d\n", toBinary(10)) // Output: 1010
	fmt.Printf("Binary of 0 is %d\n", toBinary(0))   // Output: 0

}

func toBinary(n int) int {
	if n == 0 {
		return 0
	}
	return 10 * toBinary(n / 2) + n % 2

}
