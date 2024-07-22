package main

import "fmt"

type MinHeap struct {
	array []int
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i +1
}

func right(i int) int {
	return 2*i +2
}

func (h *MinHeap) swap(first, second int) {
	h.array[first], h.array[second] = h.array[second], h.array[first]
}

func (h *MinHeap) insert(item int) {
	h.array = append(h.array, item)

	// Move node to the correct position
	h.heapifyUp(len(h.array)-1)
}

func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent((index))] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) remove() int {
	removed := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		return -1
	}

	// Move the last node to the beginning
	h.array[0] = h.array[l]

	// Slice off the last position since there is one less node in the queue
	h.array = h.array[:l]

	h.heapifyDown(0)
	return removed
}

// Heapify Down
func (h *MinHeap) heapifyDown(index int) {
	lastIndexToCheck := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
}

func main() {

}