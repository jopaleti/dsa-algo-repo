package main 

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// Pop the last two element from the stack
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform the operation and push the result back onto the stack
			var result int
			switch token {
			case "+":
				result = op1 + op2
			case "-":
				result = op1 - op2
			case "*":
				result = op1 * op2
			case "/":
				result = op1 / op2
			}
			stack = append(stack, result)
		default:
			// Convert token to integer and push unto the stack
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	
	// The final result will be the only element left in the stack
	return stack[0]
}


func main() {
	// Test cases
	tokens1 := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens1)) // Output: 9

	tokens2 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens2)) // Output: 6
}