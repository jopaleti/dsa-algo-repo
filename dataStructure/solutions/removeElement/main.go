package main

import "fmt"

// removeElement removes all occurrences of val from nums in-place
// and returns the new length of the array with elements not equal to val
func removeElement(nums []int, val int) int {
	k := 0 // Pointer to place the non-val element

	// Iterate through the array
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	// K is the number of elements not equal to val
	return k
}

func main() {
	// Test case 1
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
	
	// Test case 2
	nums = []int{0, 1, 2, 2, 3, 4, 2}
	val = 2
	k = removeElement(nums, val)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
}