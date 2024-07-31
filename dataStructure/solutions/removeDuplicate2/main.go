package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
		return len(nums)
	}

	// Pointer to write the next valid element
	write := 2

	for read := 2; read < len(nums); read++ {
		// Check if the current element is different from the element two places before
		if nums[read] != nums[write-2] {
			// Write the current element to the write position
			nums[write] = nums[read]
			write++
		}
	}

	// Return the length of the array with the valid elements
	return write
}

// Main function to test the removeDuplicates function
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println("Length of array after removing duplicates:", k)
	fmt.Println("Array after removal:", nums[:k])
}