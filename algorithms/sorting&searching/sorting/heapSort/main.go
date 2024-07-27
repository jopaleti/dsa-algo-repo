package main

import "fmt"

// Function to perform heap sort
func heapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements one by one
	for i := n-1; i > 0; i-- {
		// Swap the root element(max-element) with the last element
		arr[0], arr[i] = arr[i], arr[0]
		// Heapify the reduced heap
		heapify(arr, i, 0)
	}
}

// Function to heapify the array
func heapify(arr []int, n int, index int) {
	largest := index

	leftChild := 2*index + 1
	rightChild := 2*index + 2

	// Check if the left child exists and is greater than root
	if leftChild < n && arr[leftChild] > arr[largest] {
		largest = leftChild
	}

	// Check if the right child exists and is greater than largest so far
	if rightChild < n && arr[rightChild] > arr[largest] {
		largest = rightChild
	}

	// If largest is not root
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		// Recursively heapify the affected subtree
		heapify(arr, n, largest)
	}
}

// Main function to test the heap sort algorithm
// Main function to test the heap sort algorithm
func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", arr)
	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}