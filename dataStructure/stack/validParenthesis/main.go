package main

import (
	"fmt"
)

func isValid(s string) bool {
	// Map to keep track of matching pairs
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Stack to keep track of open brackets
	stack := []rune{}

	for _, char := range s {

		// NOTE: We use open, found to keep track of the closing parenthesis i.e '}'
		if open, found := pairs[char]; found {
			// Then, if after closing parenthesis is found, 
			// It is noted that there's no open bracket for it i.e stack
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// NOTE: if the top if statement is passed, the pop the bracket from the stack
			stack = stack[:len(stack)-1]
		} else {
			// If the character is an open bracket, push it onto the stack
			stack = append(stack, char)
		}
	}
	// If the stack is empty, all brackets are properly closed
	return len(stack) == 0
}

func main(){
	// Test cases
	tests := []struct {
		input  string
		output bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	// Run test cases
	for _, test := range tests {
		result := isValid(test.input)
		fmt.Printf("Input: %s, Output: %t\n", test.input, result)
	}
}