package main

import "fmt"

// RotateArray performs the left rotation of the array by d positions
func rotateArray(array []int, d int) {
	n := len(array)
	d = d % n // Handle cases where d >= n
	// Reverse the first d elements
	reverse(array, 0, d-1)
	// Reverse the remaining elements
	reverse(array, d, n-1)
	// Reverse the entire array
	reverse(array, 0, n-1)
}

// Reverse reverses elements in the slice from start to end index
func reverse(array []int, start, end int) {
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
}

// Example function to demonstrate the array rotation
func main() {
	// Example input
	d := 5
	a := []int{1, 2, 3, 4, 5}

	fmt.Println("Original array:", a)
	
	// Perform the rotation
	rotateArray(a, d)
	
	// Output the result
	fmt.Println("Array after rotation:", a)
}
