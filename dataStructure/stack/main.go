package main

import (
	"fmt"
)

// Mystack represents a stack data structure
type MyStack struct {
	data[]int
}

// Push inserts an element into a stack
func (s *MyStack) Push(x int) {
	s.data = append(s.data, x)
}

// IsEmpty function checks whether the stack is empty or not
func (s *MyStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Top gets the top item from the stack
func (s *MyStack) Top() int {
	
	if s.IsEmpty() {
		fmt.Println("The stack is empty")
		return -1
	}
	return s.data[len(s.data) - 1]
}

// Pop deletes an element from the stack and returns true if the operation is successful
func (s *MyStack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.data = s.data[:len(s.data)-1]

	return true
}

func main() {
	s := MyStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for i := 0; i < 4; i++ {
		if !s.IsEmpty() {
			fmt.Println(s.Top())
		}
		fmt.Println(s.Pop())
	}
}