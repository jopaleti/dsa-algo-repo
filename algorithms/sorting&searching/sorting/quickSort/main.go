package main

import "fmt"

/*
 Partition re-orders the array such that all elements less than the pivot are on 
	the left and all elements greater than the pivot are on the right and return
		the pivot index.
*/
func partition(array []int, low, high int) int {
	pivot := array[high] // Choose the rightmost element as the pivot

	i := low - 1 // Pointer for the greater element (Second pointer)

	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			// Swap array[i] with array[j]
			array[i], array[j] = array[j], array[i]
		}
	}
	// swap array[i+1] with array[high] (pivot)
	array[i+1], array[high] = array[high], array[i+1]

	return i+1
}

// Quicksort recursively sorts the array using the quick sort algorithm
func quickSort(array []int, low, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSort(array, low, pi-1)
		quickSort(array, pi+1, high)
	}
}


func main() {
    data := []int{8, 7, 2, 1, 0, 9, 6}
    
    fmt.Println("Unsorted Array")
    fmt.Println(data)
    
    quickSort(data, 0, len(data)-1)
    
    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}