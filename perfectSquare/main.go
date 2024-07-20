package main

import (
	"fmt"
	"math"
)

func numSquare(n int) int {
	// Create a slice to store the minimum number
	// of perfect squares for each number up to n
	dp := make([]int, n+1)
	fmt.Println(dp)

	// Each element is initialized to n+1,
	// which is a value larger than any possible answer.
	for i := range dp {
		dp[i] = n+1
	}
	fmt.Println(dp)

	// Base case: the least number of perfect square numbers that sum to 0 is 0
	dp[0] = 0

	// Iterate over all numbers from 1 to n
	for i := 1; i <= n; i++ {
		// Iterate over all perfect squares less than or equal to i
		for j := 1; j*j <= i; j++ {
			// Update dp[i] with minimum value
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-j*j]+1)))
			fmt.Println(i, j, dp[i])
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquare(12))
}