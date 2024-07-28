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

	// Find the smallest of the left and right children
	for l <= lastIndexToCheck {
		if l == lastIndexToCheck {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		// Swap the parent to with the smallest
		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)

			// Update the indices that we are comparing
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func main() {

	mh := &MinHeap{}
	mh.insert(35)
	mh.insert(11)
	mh.insert(23)
	mh.insert(25)
	mh.insert(17)
	mh.insert(19)
	mh.insert(30)
	mh.insert(15)
	mh.insert(8)
	mh.insert(22)
	fmt.Println(mh.array)

	removed := mh.remove()
	fmt.Printf("\nRemoved %d \n", removed)
	fmt.Println(mh.array)
}