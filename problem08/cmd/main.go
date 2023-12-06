package main

import (
	"halanproblem08/logic"
	"log"
)

func main() {
	log.Println(logic.Index_of_first_duplicate([]int{5, 17, 3, 17, 4, -1}))
	log.Println(logic.Index_of_first_duplicate([]int{}))
}
