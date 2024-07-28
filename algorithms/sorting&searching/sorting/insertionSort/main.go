package main

import "fmt"

func insertionSort(array []int) {
	for step := 1; step < len(array); step++ {
		key := array[step]
		j := step - 1

		// Compare key with each element on the left of it until an element smaller than it is found
        // For descending order, change key < array[j] to key > array[j]
		for j >= 0 && key < array[j] {
			array[j + 1] = array[j]
			j--
		}

		// Place key at the correct position
		array[j+1] = key
	}
}

func main() {
    data := []int{9, 5, 1, 4, 3}
    
    insertionSort(data)
    
    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}