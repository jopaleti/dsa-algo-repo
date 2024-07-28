#RPN -- Reverse Polish Notation
##Evaluate Reverse Polish Notation
```
###Solution
- You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
+ Evaluate the expression. Return an integer that represents the value of the expression.
```
ruby
Example 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Step-by-Step Explanation
1. Initialize a Stack:
    - Use a stack to store the operands.
    + Iterate Through Tokens:

2. For each token in the input array tokens:
    - If the token is a number, convert it to an integer and push it onto the stack.
    + If the token is an operator, pop the required operands from the stack, perform the operation, and push the result back onto the stack.
3. Return the Result:
    - After processing all tokens, the stack will contain one element, which is the result of the expression.


```go
Package strconv: This package provides functions for converting strings to different types (integers, floats, etc.) and vice versa. It is part of the Go standard library.
Function Atoi: The function Atoi stands for "ASCII to integer." It takes a string representing a number and converts it to an integer.
```