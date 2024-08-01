package main

import "fmt"

func productExceptSelf(nums []int) []int {
     n := len(nums)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0} // Edge case: no products to calculate
	}

	// Initialize result array
	result := make([]int, n)

	// Compute prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Compute suffix products and combine with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

// Main function to test the productExceptSelf function
func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println("Product except self for nums1:", productExceptSelf(nums1)) // Should output [24, 12, 8, 6]

	nums2 := []int{0, 0}
	fmt.Println("Product except self for nums2:", productExceptSelf(nums2)) // Should output [0, 0]

	nums3 := []int{1, 0, 3, 0}
	fmt.Println("Product except self for nums3:", productExceptSelf(nums3)) // Should output [0, 0, 0, 0]

	nums4 := []int{2, 3, 4, 5}
	fmt.Println("Product except self for nums4:", productExceptSelf(nums4)) // Should output [60, 40, 30, 24]
}