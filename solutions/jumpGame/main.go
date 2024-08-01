package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	farthest := 0

	for i := 0; i<len(nums); i++ {
		if i > farthest {
			return false
		}

		// i + nums[i] == means adding index of an element to it
		farthest = max(farthest, i+nums[i])
		if farthest >= len(nums)-1 {
			return true
		}
	}

	return false
}

// Main function to test the canJump function
func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println("Can jump for nums1:", canJump(nums1)) // Should output true

	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println("Can jump for nums2:", canJump(nums2)) // Should output false
}