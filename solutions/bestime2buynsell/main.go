package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        }
        profit := price - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }

    return maxProfit
}

// Main function to test the maxProfit function
func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum profit for prices1:", maxProfit(prices1)) // Should output 5

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Maximum profit for prices2:", maxProfit(prices2)) // Should output 0
}