/*
 Let’s write an algorithm to solve a common problem involving two strings:
 checking if one string is a rotation of another. 
*/

package main

import (
	"fmt"
	"strings"
)

// IsRotation checks if s2 is a rotation of s1
func isRotation(s1, s2 string) bool {
	// Checks if lengths are different
	if len(s1) != len(s2) {
		return false
	}

	// Concatenate s1 with itself
	s1s1 := s1 + s1

	// Check if s2 is a substring of s1s1
	return strings.Contains(s1s1, s2)
}

func main() {
	// Example strings
	s1 := "waterbottle"
	s2 := "erbottlewat"

	// Check if s2 is a rotation of s1
	if isRotation(s1, s2) {
		fmt.Printf("'%s' is a rotation of '%s'\n", s2, s1)
	} else {
		fmt.Printf("'%s' is not a rotation of '%s'\n", s2, s1)
	}
}