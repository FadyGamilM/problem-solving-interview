package logic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		inp [][]int
		out [][]int
	}{
		{
			inp: [][]int{{1, 2, 3}, {4, 5, 6}},
			out: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			inp: [][]int{{1}, {2}},
			out: [][]int{{1, 2}},
		},
		{
			inp: [][]int{{1}},
			out: [][]int{{1}},
		},
		{
			inp: [][]int{},
			out: [][]int{},
		},
		{
			inp: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			out: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
		},
		{
			inp: [][]int{{1, 2, 3}, {0, 0, 0}},
			out: [][]int{{1, 0}, {2, 0}, {3, 0}},
		},
		{
			inp: [][]int{{1, -2}, {-3, 4}},
			out: [][]int{{1, -3}, {-2, 4}},
		},
	}
	for _, test := range tests {
		actualOut := Transpose(test.inp)
		assert.Equal(t, test.out, actualOut)
	}
}
