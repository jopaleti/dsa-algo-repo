#QUESTION

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', 
determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 
Example 1:
Input: s = "()"
Output: true

##EXPLANATION
Define a map to store matching pairs of brackets.
Use a stack to keep track of open brackets.
Iterate through each character in the string:
If it's a closing bracket, check if it matches the top of the stack.
If it's an opening bracket, push it onto the stack.
After iterating, check if the stack is empty to determine if the string is valid.