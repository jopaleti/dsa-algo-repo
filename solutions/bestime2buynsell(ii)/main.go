package main

import "fmt"

func maxProfit(prices []int) int {
 if len(prices) == 0 {
	return 0
 }

 maxProfit := 0
 for i := 1; i < len(prices); i++ {
	if prices[i] > prices[i-1] {
		maxProfit += prices[i] - prices[i-1]
	}
 }
 return maxProfit
}

// Main function to test the maxProfit function
func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum profit for prices1:", maxProfit(prices1)) // Should output 7

	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Maximum profit for prices2:", maxProfit(prices2)) // Should output 4

	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println("Maximum profit for prices3:", maxProfit(prices3)) // Should output 0
}