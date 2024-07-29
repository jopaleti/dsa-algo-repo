package main

import "fmt"

// CountingSort sorts the array based on the digit represented by "Place"
func countingSort(array []int, place int) {
	size := len(array)
	output := make([]int, size)
	count := make([]int, 10)

	// Calculate count of elements
	for _, value := range array {
		// NOTE: Value of an array represent index of count whose value is 
		//	the number of occurence of index in the array
		index := value / place
		count[index % 10]++
	}

	// Calculate the cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Place element in sorted order
	for i := size - 1; i >= 0; i-- {
		index := array[i] / place
		output[count[index % 10] - 1] = array[i]
		count[index % 10]--
	}

	// Copy sorted element into the original array
	for i := 0; i < size; i++ {
		array[i] = output[i]
	}
}

// RadixSort sorts the array using radix sort algorithm
func radixSort(array []int) {
	if len(array) == 0 {
		return
	}

	// Get the maximum element
	maxElement := array[0]
	for _, value := range array {
		if value > maxElement {
			maxElement = value
		}
	}

	// Apply countingSort to sort elements based on place value
	place := 1
	for maxElement / place > 0 {
		countingSort(array, place)
		place *= 10
	}
}

func main() {
    data := []int{121, 432, 564, 23, 1, 45, 788}

    radixSort(data)

    fmt.Println("Sorted Array:")
    fmt.Println(data)
}