package main

import (
	"fmt"
)

// Doubly linked list node
type node struct {
	value int
	next *node
	prev *node
}

// String method for the node to print its value
func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Defining the doubly linked list struct
type doublyLinkedList struct {
	head *node
	tail *node
	len int
}

// Inserting a new node at the end of the list
func (dll *doublyLinkedList) add(value int) {
	newNode := &node{value: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.len++
}

// Removing a node from the list
func (dll *doublyLinkedList) remove(value int) {
	if dll.head == nil {
		return
	}	

	current := dll.head
	for current != nil {
		if current.value == value {
			// 1 check if current is the head
			// 2 check if current is the tail
			// 3 check if current is in the middle
			if current == dll.head {
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				} else {
					dll.tail = nil
				}
			} else if current == dll.tail {
				dll.tail = current.prev
				dll.tail.next = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			dll.len--
			return
		}
		current = current.next
	}
}

// Retrieve a node from the list
func (dll doublyLinkedList) get(value int) *node {
	for iterator := dll.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

// String method for the doubly linked list to print its nodes
