package main

import "fmt"

func intToRoman(num int) string {
	// Define the mapping of roman numeral symbols
	vals := []struct{
		Value int
		Symbol string
	} {
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// Initialize an empty result string
	result := ""

	// Iterate over values from largest to smallest
	for _, pair := range vals {
		for num >= pair.Value {
			result += pair.Symbol
			num -= pair.Value
		}
	}

	return result
}

func main() {
	num := 3749
	fmt.Println("Integer to Roman:", intToRoman(num)) // Should output MMMDCCXLIX
}