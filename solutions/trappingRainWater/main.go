package main

import "fmt"


func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	trappedWater := 0

	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				trappedWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				trappedWater += rightMax - height[right]
			}
			right--
		}
	}

	return trappedWater
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Trapped Water:", trap(height)) // Should output 6
}