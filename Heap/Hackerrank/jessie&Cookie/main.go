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


