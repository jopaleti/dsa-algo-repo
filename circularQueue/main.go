package main

import "fmt"

type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
	capacity int
}

// Constructor(k int) MyCircularQueue: Initializes the queue with a specified size k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: make([]int, k),
		head: -1,
		tail: -1,
		size: 0,
		capacity: k,
	}
}

// Enqueue: - Adding element to the rare or back of the list
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	// The tail pointer is updated to point to the next position in the circular buffer.
	this.tail = (this.tail + 1) % this.capacity
	this.data[this.tail] = value
	this.size++
	return true
}

// Dequeue: Removing element from the front of the queue
func (this *MyCircularQueue) DeQueue(value int) bool {
	if this.IsEmpty() {
		return false
	}
	// If the head pointer is equal to the tail pointer, 
	// it means the queue has only one element.
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		// If the queue has more than one element, the head pointer is updated 
		// to point to the next position in the circular buffer. 
		this.head = (this.head + 1) % this.capacity
	}
	this.size--
	return true
}

// Checking the front element
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

// Checking for the rear or last element
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

// Checking if our circular list is empty
// NOTE: If the queue is empty, size must be zero(0)
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

// Checking if our queue is full
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

func main() {
	 // Initialize the circular queue with a capacity of 3
    obj := Constructor(3)
    
    // Test EnQueue
    fmt.Println(obj.EnQueue(1)) // true
    fmt.Println(obj.EnQueue(2)) // true
    fmt.Println(obj.EnQueue(3)) // true
    fmt.Println(obj.EnQueue(4)) // false (queue is full)

	// Test Rear
    fmt.Println(obj.Rear()) // 3
    
    // Test IsFull
    fmt.Println(obj.IsFull()) // true
    
    // Test DeQueue
    fmt.Println(obj.DeQueue(4)) // true

	// Test EnQueue after DeQueue
    fmt.Println(obj.EnQueue(4)) // true
    
    // Test Rear after EnQueue
    fmt.Println(obj.Rear()) // 4
    
    // Test Front
    fmt.Println(obj.Front()) // 2
    
    // Test IsEmpty
    fmt.Println(obj.IsEmpty()) // false
}