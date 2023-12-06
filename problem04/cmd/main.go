package main

import "log"

type operation func(x int) int

func main() {
	increment := operation(func(x int) int {
		return x + 1
	})
	square := operation(func(x int) int {
		return x * x
	})

	h := composite(square, increment)
	log.Println(h(6))
	log.Println(h(7))
	log.Println(h(8))
	log.Println(h(0))

}

func composite(arg1, arg2 operation) func(x int) int {
	return func(x int) int {
		return arg1(arg2(x))
	}
}
