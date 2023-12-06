package logic

import (
	"unicode"
)

/*
Edge cases I will assume
 1. If the input string is an empty string ➜  return true
 2. If the input string is one character ➜  return true
 3. The spaces and punctuations are ignores
 4. small letters are considered the same as capital letters
*/
func Palindrome(inp string) bool {
	data := []rune(inp)
	ignored := setupIgnoredChars()
	// ingore the spaces by trimming them to get a clean []rune
	data = ignoreSpaces(data)

	// edge case where its an empty string
	if len(data) == 0 || len(data) == 1 {
		return true
	}

	// define the two pointers
	start := 0
	end := len(data) - 1

	for start <= end {
		// log.Printf("now test the start -> %v and end -> %v\n", string(data[start]), string(data[end]))
		// if we have a string with odd length (5 chars) / (1 char)
		if start == end {
			return true
		} else {
			_, start_ok := ignored[data[start]]
			_, end_ok := ignored[data[end]]
			if start_ok && end_ok {
				start++
				end--
				continue
			}
			if start_ok {
				start++
				continue
			}
			if end_ok {
				end--
				continue
			}
			if unicode.ToLower(data[start]) == unicode.ToLower(data[end]) {
				// move start 1 step forward and move end 1 step backward
				start++
				end--
				continue
			} else {
				// log.Printf("the start which failed : [%v] %v \n", string(data[start]), unicode.ToLower(data[start]))
				// log.Printf("the end which failed : [%v] %v \n", string(data[end]), unicode.ToLower(data[end]))
				return false
			}
		}
	}

	return true
}

func ignoreSpaces(data []rune) []rune {
	output := []rune{}
	for _, c := range data {
		if c != ' ' {
			output = append(output, c)
		}
	}
	return output
}

func setupIgnoredChars() map[rune]bool {
	ignored := map[rune]bool{}
	ignored[rune('!')] = true
	ignored[rune('@')] = true
	ignored[rune('#')] = true
	ignored[rune('$')] = true
	ignored[rune('%')] = true
	ignored[rune('^')] = true
	ignored[rune('&')] = true
	ignored[rune('*')] = true
	ignored[rune('(')] = true
	ignored[rune(')')] = true
	ignored[rune('-')] = true
	ignored[rune('=')] = true
	ignored[rune('+')] = true
	ignored[rune('-')] = true
	ignored[rune('_')] = true
	ignored[rune('.')] = true
	ignored[rune(',')] = true
	ignored[rune('~')] = true
	ignored[rune('?')] = true
	ignored[rune('/')] = true
	ignored[rune(']')] = true
	ignored[rune('[')] = true
	ignored[rune('{')] = true
	ignored[rune('}')] = true
	ignored[rune('<')] = true
	ignored[rune('>')] = true
	ignored[rune(':')] = true
	ignored[rune(';')] = true
	ignored[rune('"')] = true
	return ignored
}
