package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfFirstDuplicate(t *testing.T) {
	testCases := []struct {
		inp         []int
		expectedOut int
	}{
		{inp: []int{}, expectedOut: -1},
		{inp: []int{1, 1, 2}, expectedOut: 1},
		{inp: []int{1, 2, 3, 4, 5, 6, 5}, expectedOut: 6},
	}

	for _, testCase := range testCases {
		actualOut := Index_of_first_duplicate(testCase.inp)
		assert.Equal(t, testCase.expectedOut, actualOut)
	}
}
