package main

import (
	"fmt"
)

// Function to reverse a portion of the array in-place
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Function to rotate the array to the right by k steps
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// Normalize k to be within the bounds of the array
	k = k % n
	if k == 0 {
		return
	}

	// Reverse the entire array
	reverse(nums, 0, n-1)
	// Reverse the first k elements
	reverse(nums, 0, k-1)
	// Reverse the remaining n-k elements
	reverse(nums, k, n-1)
}

// Main function to test the rotate function
func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println("Rotated array:", nums1) // Should output [5, 6, 7, 1, 2, 3, 4]

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println("Rotated array:", nums2) // Should output [3, 99, -1, -100]
}
