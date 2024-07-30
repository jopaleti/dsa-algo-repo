package main

import (
	"fmt"
)

// Function to find the minimum and maximum values in a slice
func findMinMax(slice []int) (int, int) {
	if len(slice) == 0 {
		fmt.Println("The slice is empty")
		return 0, 0
	}

	// Initialize min and max with the first element of the slice
	min := slice[0]
	max := slice[0]

	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

// Main function to test the findMinMax function
func main() {
	slice := []int{3, 5, 1, 8, 7, 2}
	fmt.Println("Array:", slice)
	min, max := findMinMax(slice)
	fmt.Println("Min value:", min)
	fmt.Println("Max value:", max)
}