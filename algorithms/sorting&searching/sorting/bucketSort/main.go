package main

import "fmt"

// Function to perform bucket sort
func bucketSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// Finding the maximum value to determine the range of values
	maxValue := arr[0]
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
	}

	// Create Buckets
	numBuckets := maxValue + 1
	buckets := make([][]int, numBuckets)

	// Distribute elements into buckets
	// NOTE: The value in an array is the index and still the value in the bucket
	for _, value := range arr {
		buckets[value] = append(buckets[value], value)
	}

	// Concatenate buckets into the original array
	index := 0
	for _, bucket := range buckets {
		for _, value := range bucket {
			arr[index] = value
			index++
		}
	}
}


// Main function to test the bucket sort algorithm
func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Unsorted array:", arr)
	bucketSort(arr)
	fmt.Println("Sorted array:", arr)
}