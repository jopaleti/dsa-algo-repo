package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt32},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	min := this.minStack[len(this.minStack)-1]

	if val < min {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// Test cases
func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(5)
	obj.Push(2)
	obj.Push(1)
	fmt.Println(obj.Top())    // Should print 1
	fmt.Println(obj.GetMin()) // Should print 1
	obj.Pop()
	fmt.Println(obj.Top())    // Should print 2
	fmt.Println(obj.GetMin()) // Should print 2
}