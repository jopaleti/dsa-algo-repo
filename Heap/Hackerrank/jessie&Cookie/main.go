// Date: 2024/07/26
// Author: Opaleti Oluwatobi
// Description: This is a solution to the problem on Hackerrank called Jessie and Cookies
// Check docs.md for more information about the problem and the solution

package main

import "fmt"

// Defining the heap struct
type MinHeap struct {
	elements []int
}

// Parent declaration (Function)
func parent(i int) int {
	return i - 1 / 2
}

// Left declaration (Function)
func left(i int) int {
	return 2*i + 1
}

// Right declaration (Function)
func right(i int) int {
	return 2*i + 2
}

// SWAP function declaration
func (h *MinHeap) swap(first, second int) {
	h.elements[first], h.elements[second] = h.elements[second], h.elements[first]
}

// Insert function declaration
func (h *MinHeap) Insert(x int) {
	h.elements = append(h.elements, x)
	h.heapifyUp(len(h.elements)-1)
}

// Delete Function
func (h *MinHeap) Delete(x int) {
	for i, v := range h.elements {
		if v == x {
			h.elements[i] = h.elements[len(h.elements) - 1]
			h.elements = h.elements[:len(h.elements)-1]
			h.heapifyDown(i)
			break
		}
	}
}

// GetMin function
func (h *MinHeap) GetMin() int {
	if len(h.elements) == 0 {
		return -1 // Assuming non-negative integers in the heap
	}
	return h.elements[0]
}

// DeleteMin function
func (h *MinHeap) DeleteMin() int {
	if len(h.elements) == 0 {
		return -1
	}
	min := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.heapifyDown(0)
	return min
}

// HeapifyUp Algorithm
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		if h.elements[index] >= h.elements[parent(index)] {
			break
		}
		h.swap(index, parent(index))
		// Update the indices
		index = parent(index)
	}
}

// HeapifyDown Algorithm
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.elements) - 1

	for {
		leftChild := left(index)
		rightChild := right(index)
		smallest := index

		if leftChild <= lastIndex && h.elements[leftChild] < h.elements[smallest] {
			smallest = leftChild
		}
		if rightChild <= lastIndex && h.elements[rightChild] < h.elements[smallest] {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		h.swap(index, smallest)
		// Updates the indices
		index = smallest
	}
}


// Cookies function
func cookies(k int, A []int) int {
	h := &MinHeap{}

	for _, sweetness := range A {
		h.Insert(sweetness)
	}

	operations := 0
	for len(h.elements) > 1 {
		// Check if the smallest element is already greater than or equal to k
		if h.elements[0] >= k {
			return operations
		}

		// Remove the two least sweet cookies
		first := h.DeleteMin()
		second := h.DeleteMin()

		// Combine them into a new cookie
		newCookie := first + 2*second

		// Insert the new cookie back into the heap
		h.Insert(newCookie)

		// Increment the operation count
		operations++
	}

	// If the least sweet cookie is still less than k, return -1
	if h.elements[0] < k {
		return -1
	}

	return operations
}

// TEST SAMPLES
func main() {
	k := 7
	A := []int{1, 2, 3, 9, 10, 12}
	fmt.Println(cookies(k, A)) // Output: 2
}