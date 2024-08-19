package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// Get the value of the current roman numeral
		current := romanMap[rune(s[i])]
		// Get the value of the next roman numeral, if available
		var next int
		if i+1 < n {
			next = romanMap[rune(s[i+1])]
		}

		// Determine whether to add or subtract the current value
		if current < next {
			total -= current
		} else {
			total += current
		}
	}

	return total
}

func main() {
	example := "MCMXCIV"
	fmt.Println("Roman to Integer:", romanToInt(example)) // Should output 1994
}