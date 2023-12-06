package logic

// receives slice of integeres and returns the index of the first duplicate, if there is no it returns -1
/*
	===> Time Complexity is O(n) in the most bad worst depressed case :( the complexity will be O(n*n) because the hash collosion but its too rare
	===> Space Complexity is O(n) in worst case all items are unqiue
*/
func Index_of_first_duplicate(inp []int) int {
	seenBefore := make(map[int]bool, 0)
	for idx, item := range inp {
		if seenBefore[item] {
			return idx
		}
		seenBefore[item] = true
	}
	return -1
}
