package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)

	if n == 0 {
		return -1
	}

	totalGas, totalCost := 0, 0
	currentTank, startIndex := 0, 0

	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentTank += gas[i] - cost[i]

		if currentTank < 0 {
			// Reset the currentTank and startIndex
			startIndex = i + 1
			currentTank = 0
		}
	}

	// Check if the total gas is enough for the total cost
	if totalGas < totalCost {
		return -1
	}

	return startIndex
}

// Main function to test the canCompleteCircuit function
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println("Starting gas station index:", canCompleteCircuit(gas, cost)) // Should output 3
}