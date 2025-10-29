package main

import (
	"fmt"
	"Task3/stats"
)

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	largest := stats.Largest(numbers)
	smallest := stats.Smallest(numbers)
	average := stats.Average(numbers)
	fmt.Printf("Largest: %d\n", largest)
	fmt.Printf("Smallest: %d\n", smallest)
	fmt.Printf("Average: %.2f\n", average)
}
