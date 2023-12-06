package main

import (
	"halanproblem05/logic"
	"log"
)

func main() {
	// words := []string{"apples", "oranges", "apples", "flowers"}
	words := []string{"apples", "apples"}
	output := logic.GetUniqueWords_optimized(words)
	log.Println(output)
}
