package main

import "fmt"

func jump(nums []int) int {
    n := len(nums)
    if n<=1 {
        return 0
    }
    jumps := 0
    currentEnd := 0
    farthest := 0 

    for i:=0; i<=n-1; i++ {
        farthest = max(farthest, i+nums[i])
        if i == currentEnd {
            jumps++
            currentEnd = farthest
            if currentEnd >= n-1 {
                break
            }
        }
    }

    return jumps
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Main function to test the jump function
func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Printf("Minimum jumps for nums1: %d\n", jump(nums1)) // Should output 2

	nums2 := []int{2, 3, 0, 1, 4}
	fmt.Printf("Minimum jumps for nums2: %d\n", jump(nums2)) // Should output 2

	nums3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Minimum jumps for nums3: %d\n", jump(nums3)) // Should output 3
}
