# Problem 01
A palindrome is a word or sequence of characters which reads the same backward or
forward. For example the words: level, anna, noon, rotator are all palindromes.
Write a function palindrom that accepts a string as an argument and returns a boolean
indicating whether the input is a palindrome or not, for example: <br>
```go
palindrome("anna")  // returns True
palindrome("apple") //  returns False
```

# My Solution & Assumptions 
	Edge cases I will assume
		1. If the input string is an empty string   ➜  return true
		2. If the input string is one character     ➜  return true
		3. The spaces and punctuations are ignores in the comparison 
		4. small letters are considered the same as capital letters

