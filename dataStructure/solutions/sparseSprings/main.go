// Sparse strings typically refer to strings where most of the elements are 
// the same or have a common pattern, allowing them to be represented more efficiently.

package main

import (
	"fmt"
	"strings"
)

// CountOccurrences counts how many times the query string appears in the text
func countOccurrences(text, query string) int {
	count := 0
	for {
		index := strings.Index(text, query)
		if index == -1 {
			break
		}
		count++
		// Move past the current occurrences 
		text = text[index + len(query):]
	}
	return count
}

// Main function to demonstrate the usage
func main() {
	// Example text and queries
	text := "ababcabcab"
	queries := []string{"ab", "abc", "bca"}

	// Iterate over each query and count occurrences in the text
	for _, query := range queries {
		count := countOccurrences(text, query)
		fmt.Printf("Query '%s' appears %d times in the text.\n", query, count)
	}
}