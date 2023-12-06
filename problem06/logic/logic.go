package logic

func Transpose(inp [][]int) [][]int {
	if len(inp) == 0 {
		return [][]int{}
	}
	inp_rows := len(inp)
	inp_cols := len(inp[0])
	// log.Printf("rows %v cols %v\n", inp_rows, inp_cols)

	out := make([][]int, inp_cols) // the cols of input are rows of output

	// now append the inner slices
	for i := 0; i < inp_cols; i++ {
		out[i] = make([]int, inp_rows)
	}

	// log.Println(len(out))
	// log.Println(len(out[0]))

	for i := 0; i < inp_rows; i++ {
		for j := 0; j < inp_cols; j++ {
			out[j][i] = inp[i][j]
		}
	}

	return out
}
