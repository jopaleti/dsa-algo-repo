package main

import (
	"fmt"
)

// Function to calculate the sum of elements in a slice
func sumArray(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

// Main function to test the sumArray function
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Array:", slice)
	sum := sumArray(slice)
	fmt.Println("Sum of elements:", sum)
}