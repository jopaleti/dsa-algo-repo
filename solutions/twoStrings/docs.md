#Problem Statement:
Given two strings s1 and s2, determine if s2 is a rotation of s1. For example:
"waterbottle" is a rotation of "erbottlewat".
"hello" is not a rotation of "world".

##Solution Approach
```
1. Concatenate s1 with itself: This allows us to use a substring search to check
 if s2 exists within this concatenated string. If s2 is a rotation of s1, 
 it should appear in s1 + s1.
2. Check length: Ensure both strings are of the same length. 
```

##Algorithm
```
1. Check lengths: If s1 and s2 are not of the same length, return false.
2. Concatenate s1: Create a new string s1s1 by concatenating s1 with itself.
3. Search for s2: Check if s2 is a substring of s1s1.
4. Return the result: If s2 is found in s1s1, return true; otherwise, return false.
```
