# Explanation
1. Mapping:
- romanMap is a dictionary that maps each Roman numeral character to its integer value.
2. Iteration:
- For each character in the string, compare its value with the value of the next character.
- If the current value is less than the next value, subtract the current value from the total (this handles the special cases like IV and IX).
- Otherwise, add the current value to the total.
3. Edge Case Handling:
- The condition if current < next ensures that special subtraction cases are handled correctly.

### Time Complexity
1. Time Complexity: O(n) – where n is the length of the input string.
2. Space Complexity: O(1) – as the space used does not grow with input size.