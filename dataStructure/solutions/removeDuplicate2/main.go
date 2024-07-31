package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    write := 2
    for read := 2; read<len(nums); read++ {
        if nums[read] != nums[write-2] {
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