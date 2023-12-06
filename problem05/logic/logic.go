package logic

// import (
// "hash/fnv"
// )

// func hash(inp string) uint64 {
// 	h := fnv.New64a()
// 	h.Write([]byte(inp))
// 	return h.Sum64()
// }

func GetUniqueWords(words []string) []string {
	items := make(map[string]int, 0)
	// checkMap := make([]uint64, 0)
	output := make([]string, 0)
	// we store all of them
	for _, word := range words {
		items[word]++
	}
	for word, freq := range items {
		if freq == 1 {
			output = append(output, word)
		}
	}
	return output
}

// instead of looping through all words even the repeated ones again, we just loop through them once
// i know that we drop the constants in the big O complexity but with 1 billion slice, i think looping through 2 billion is not as looping through 1 billion
func GetUniqueWords_optimized(words []string) []string {
	items := make(map[string]bool, 0)
	// checkMap := make([]uint64, 0)
	output := make([]string, 0)
	// in this for loop i used the hash map so instead of storing all words with their frequency and then looping through the entire lsit of words, i only loop through the unique ones (the ones i need)
	for _, word := range words {
		_, ok := items[word]
		if !ok {
			items[word] = true
		} else {
			delete(items, word)
		}
	}
	for unique, _ := range items {
		output = append(output, unique)
	}
	return output
}
