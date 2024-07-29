package main

import "fmt"

// Merge function merges two slices into one sorted slice.
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge the two slices while there are elements in both slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements from the right slice
	for j < len(right) {
		result = append(result, right[i])
		j++
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle element
	mid := len(arr) / 2

	// Recursively split and sort the slices 
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted slices
	return merge(left, right)
}

func main() {
	data := []int{4, 2, 2, 8, 3, 3, 1}

    sortedArr := mergeSort(data)

    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(sortedArr)
}