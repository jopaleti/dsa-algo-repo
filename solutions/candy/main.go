package main

import "fmt"

/*
	NOTE: The solution to the algorithm is obtained through:-
		TWO PASS APPROACH
*/

// Function to compute minimum number of candies required
func candy(ratings []int) int {
	n := len(ratings)

	if n == 0 {
		return 0
	}

	// Initialize candies array and assign 1 candy to each initially
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// Left to Right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right to Left
	for i := n-2; i>=0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Sum up the candies
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function to test the candy function
func main() {
	ratings := []int{1, 0, 2}
	fmt.Println("Minimum number of candies required:", candy(ratings)) // Should output 5
}