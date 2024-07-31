package main

import "fmt"

func majorityElement(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    majority := make(map[int]int)
    for _, value := range nums {
        majority[value]++
    }
    var majorityElement int
    maxcount := 0
    for key, count := range majority {
        if count > maxcount {
            maxcount = count
            majorityElement = key
        }
    }
    // Return majority element if its count is greater than the half length of an array
    if maxcount > len(nums)/2 {
        return majorityElement
    }
    return majorityElement
}

// Main function to test the majorityElement function
func main() {
	nums1 := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	fmt.Println("Majority element in nums1:", majorityElement(nums1)) // Should output 4

	nums2 := []int{3, 3, 4, 2, 4, 4, 2, 4}
	fmt.Println("Majority element in nums2:", majorityElement(nums2)) // Should output -1 (no majority element)
}