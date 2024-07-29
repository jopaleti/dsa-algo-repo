package main

import "fmt"

// CountingSort sorts an array of integers using the counting sort algorithm
func countingSort(array []int) {
	size := len(array)
	if size == 0 {
		return 
	}

	// Finding the maximum value in the array to determine the range of count array
	maxVal := array[0]
	for _, value := range array {
		if value > maxVal {
			maxVal = value
		}
	}

	// Initialize count array
	count := make([]int, maxVal + 1)
	output := make([]int, size)

	// Store the count of each element in the count array
	for _, value := range array {
		count[value]++
	}

	// Store the cumulative count
	for i := 1; i<len(count); i++ {
		count[i] += count[i-1]
	}

	// Place the elements in the output array
	for i := size - 1; i >= 0; i-- {
		output[count[array[i]]-1] = array[i]
		count[array[i]]--
	}

	// Copy the sorted array into the original array
	copy(array, output)

}

// Test samples 
func main() {
    data := []int{4, 2, 2, 8, 3, 3, 1}

    countingSort(data)

    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}