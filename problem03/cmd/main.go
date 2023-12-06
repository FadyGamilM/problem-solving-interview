package main

import (
	"halanproblem03/logic"
	"log"
)

func main() {
	inp := "aaaaaaaaaabbbaxxxxyyyzyx"
	sw := logic.NewSlidingWindow()
	encoded := sw.RunLengthEncode(inp)
	log.Printf("ENCODING ➜ %v\n", encoded)
	decoded := sw.RunLengthDecode(encoded)
	log.Printf("DECODING ➜ %v\n", decoded)
	if inp == decoded {
		log.Println("the input is the same as output ")
	}
	sw.Reset()
	encoded = sw.RunLengthEncode("xx")
	log.Println(encoded)

	sw.Reset()
	encoded = sw.RunLengthEncode("5")
	log.Println(encoded)
}
