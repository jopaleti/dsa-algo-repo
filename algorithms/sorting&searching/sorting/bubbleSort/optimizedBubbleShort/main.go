package main

import "fmt"

/* bubbleSort sorts an array of integers in ascending order
using an optimized bubble sort algorithm
*/

func bubbleSort(array []int) {
	n := len(array)

	// Outer loop to traverse through the array
	for i := 0; i < n; i++ {
		/*
			SWAPPED is used to keep track of whether
			 any elements were swapped during the inner loop.

			 NOTE: It is put inside the outer loop to reset after 
			       each iteration
		*/
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				// Swap elements if they are in the wrong order
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		 // If no elements were swapped during an entire pass, 
		//  the array is already sorted, so the loop breaks early.
		 if !swapped {
			break
		 }
	}
}


func main() {
    data := []int{-2, 45, 0, 11, -9}
    
    bubbleSort(data)
    
    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}
