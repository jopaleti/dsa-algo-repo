package main

import (
	"fmt"
)

// Function to insert element into a slice at a specific index
func insertElement(slice []int, index int, value int) []int {
	if index < 0 || index > len(slice) {
		// Index out of range, return the original slice
		fmt.Println("Index out of range")
		return slice
	}

	// Create a new slice with one more capacity
	newSlice := make([]int, len(slice)+1)

	// Copy elements before the index
	copy(newSlice, slice[:index])

	// Insert the new element
	newSlice[index] = value

	// Copy elements after the index
	copy(newSlice[index+1:], slice[index:])

	return newSlice
}

// Main function to test the insertElement function
func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", slice)
	
	indexToInsert := 2
	valueToInsert := 25
	slice = insertElement(slice, indexToInsert, valueToInsert)
	
	fmt.Println("Slice after insertion:", slice)
}