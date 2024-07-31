package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    k := 0

    for i:=1; i<len(nums); i++ {
        if nums[i] != nums[k] {
            k++
            nums[k] = nums[i]
        }
    }

    return k+1
}

func main() {
	// Test case 1
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
	
	// Test case 2
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 4}
	k = removeDuplicates(nums)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
	
	// Test case 3
	nums = []int{1, 1, 1, 1, 1}
	k = removeDuplicates(nums)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
	
	// Test case 4
	nums = []int{}
	k = removeDuplicates(nums)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
}