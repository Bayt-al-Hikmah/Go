package main

import (
	"fmt"
	"math"
)

func SafeSqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("cannot calculate sqrt of negative number: %f", num)
	}
	return math.Sqrt(num), nil
}

func main() {
	result, err := SafeSqrt(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result (positive):", result)
	}

	result, err = SafeSqrt(-9)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result (negative):", result)
	}
}