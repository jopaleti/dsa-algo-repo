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
func (thisqueue *MyCircularQueue) EnQueue(value int) bool {
	if thisqueue.IsFull() {
		return false
	}
	if thisqueue.IsEmpty() {
		thisqueue.head = 0
	}
	// The tail pointer is updated to point to the next position in the circular buffer.
	thisqueue.tail = (thisqueue.tail + 1) % thisqueue.capacity
	thisqueue.data[thisqueue.tail] = value
	thisqueue.size++
	return true
}

// Dequeue: Removing element from the front of the queue
func (thisqueue *MyCircularQueue) DeQueue(value int) bool {
	if thisqueue.IsEmpty() {
		return false
	}
	// If the head pointer is equal to the tail pointer, 
	// it means the queue has only one element.
	if thisqueue.head == thisqueue.tail {
		thisqueue.head = -1
		thisqueue.tail = -1
	} else {
		// If the queue has more than one element, the head pointer is updated 
		// to point to the next position in the circular buffer. 
		thisqueue.head = (thisqueue.head + 1) % thisqueue.capacity
	}
	thisqueue.size--
	return true
}

// Checking the front element
func (thisqueue *MyCircularQueue) Front() int {
	if thisqueue.IsEmpty() {
		return -1
	}
	return thisqueue.data[thisqueue.head]
}

// Checking for the rear or last element
func (thisqueue *MyCircularQueue) Rear() int {
	if thisqueue.IsEmpty() {
		return -1
	}
	return thisqueue.data[thisqueue.tail]
}

// Checking if our circular list is empty
// NOTE: If the queue is empty, size must be zero(0)
func (thisqueue *MyCircularQueue) IsEmpty() bool {
	return thisqueue.size == 0
}

// Checking if our queue is full
func (thisqueue *MyCircularQueue) IsFull() bool {
	return thisqueue.size == thisqueue.capacity
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