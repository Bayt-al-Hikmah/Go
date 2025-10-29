package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a slice of colors
	colors := []string{"red", "green", "blue"}
	fmt.Println("First element:", colors[0])
	fmt.Println("Last element:", colors[len(colors)-1])

	colors = append(colors, "yellow")
	colors = append([]string{"black"}, colors...)
	fmt.Println("After adding elements:", colors)
	
	lastColor := colors[len(colors)-1]
	colors = colors[:len(colors)-1]
	fmt.Println("Removed last element:", lastColor)
	fmt.Println("After removing last element:", colors)

	containsGreen := false
	for _, color := range colors {
		if color == "green" {
			containsGreen = true
			break
		}
	}
	fmt.Println("Contains 'green':", containsGreen)

	// Number slice
	numbers := []int{5, 3, 8, 1}
	numbers = append(numbers[:1], numbers[2:]...) 
	fmt.Println("After removing second item:", numbers)
	
	numbers = append(numbers[:2], append([]int{25}, numbers[2:]...)...)
	fmt.Println("After inserting 25 at index 2:", numbers)

	// Sort the slice
	sort.Ints(numbers)
	fmt.Println("Sorted slice:", numbers)
}