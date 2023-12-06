package main

import (
	"halanproblem01/logic"
	"log"
)

func main() {
	out := logic.Palindrome("Dog doo? Good God!")
	log.Println("result is : ", out)

	out = logic.Palindrome("åbba")
	log.Println("result is : ", out)
}
