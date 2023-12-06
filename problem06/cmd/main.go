package main

import (
	"fmt"
	"halanproblem06/logic"
)

func main() {
	inp := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// inp_rows := len(inp)
	// inp_cols := len(inp[0])
	// // log.Printf("rows %v cols %v\n", inp_rows, inp_cols)

	// out := make([][]int, inp_cols) // the cols of input are rows of output

	// // now append the inner slices
	// for i := 0; i < inp_cols; i++ {
	// 	out[i] = make([]int, inp_rows)
	// }

	// // log.Println(len(out))
	// // log.Println(len(out[0]))

	// for i := 0; i < inp_rows; i++ {
	// 	for j := 0; j < inp_cols; j++ {
	// 		out[j][i] = inp[i][j]
	// 	}
	// }

	out := logic.Transpose(inp)

	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out[0]); j++ {
			fmt.Printf(" %v ", out[i][j])
		}
		fmt.Println()
	}
	inp = [][]int{{1}, {2}}
	out = logic.Transpose(inp)
	fmt.Println()
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out[0]); j++ {
			fmt.Printf(" %v ", out[i][j])
		}
		fmt.Println()
	}
}
