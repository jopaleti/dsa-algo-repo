package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	/*
	Create an array answer of the same length as temperatures, initialized with zeros.
	This will store the result.
	*/
	answer := make([]int, n)
	stack := []int{}	// Stack to keep track of indices

	// Loop through each temperature in the temperatures array.
	for i := 0; i < n; i++ {
		// Check if current temperature is greater than the temperature at indices in stack
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// Pop the index from the stack
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Update the answer for the popped index
			answer[prevIndex] = i - prevIndex
		}
		// Push the current index onto the stack
		stack = append(stack, i)
	}
	return answer   
}

func main() {
	// Test case
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temperatures)
	fmt.Println(result) // Output: [1, 1, 4, 2, 1, 1, 0, 0]
}