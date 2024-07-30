package main

import (
	"fmt"
)

// Function to delete an element from a slice
func deleteElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// Index out of range, return the original slice
		fmt.Println("Index out of range")
		return slice
	}

	// Remove element by slicing and concatenating
	return append(slice[:index], slice[index+1:]...)
}

// Main function to test the deleteElement function
func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", slice)
	
	indexToDelete := 2
	slice = deleteElement(slice, indexToDelete)
	
	fmt.Println("Slice after deletion:", slice)
}