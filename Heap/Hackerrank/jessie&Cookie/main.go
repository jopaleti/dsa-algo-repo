package main

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
			smallest = rightChild
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


func main() {
	
}