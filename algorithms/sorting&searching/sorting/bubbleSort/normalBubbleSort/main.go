package main

import "fmt"

/* BUBBLESORT: It is a simple sorting algorithm that repeatedly
	 steps through the list, compares adjacent elements and 
	 swaps them if they are in the wrong order.
*/

/*BubbleSort sorts an array of integers in ascending order
	  using the bubble sort algorithm. */

func bubbleSort(array []int) {
	n := len(array)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				// Swap the elements if they are in the wrong order
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}

	bubbleSort(data)
	fmt.Println("Sorted Array in Ascending Order:", data)
}