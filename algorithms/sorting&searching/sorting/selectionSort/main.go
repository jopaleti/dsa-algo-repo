package main

import "fmt"

/* SelectionSort sorts an array of integers in
   ascending order using selection sort algorithm */
func selectionSort(array []int) {
	size := len(array)
	for step := 0; step < size; step++ {
		minIdx := step

		for i := step + 1; i < size; i++ {
			if array[i] < array[minIdx] {
				minIdx = i
			}
		}

		// Swap the found minimum element with the element at the current step
		array[step], array[minIdx] = array[minIdx], array[step]
	}
}

func main() {
    data := []int{-2, 45, 0, 11, -9}
    
    selectionSort(data)
    
    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}