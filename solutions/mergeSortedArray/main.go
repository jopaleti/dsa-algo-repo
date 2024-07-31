package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// Pointers for nums1, nums2, and the ends of nums1
	i := m -1
	j := n -1
	k := m+n -1

	// Merging nums1 and nums2 from the end
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// Copy remaining elements from num2, if any
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func main() {
	// Test case 1
    nums1 := []int{1, 2, 3, 0, 0, 0}
    m := 3
    nums2 := []int{2, 5, 6}
    n := 3
    
    fmt.Println("Before merge:")
    fmt.Println("nums1:", nums1)
    fmt.Println("nums2:", nums2)
    
    merge(nums1, m, nums2, n)
    
    fmt.Println("After merge:")
    fmt.Println("nums1:", nums1)
}