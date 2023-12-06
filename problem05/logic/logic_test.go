package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUniqueWords(t *testing.T) {
	tests := []struct {
		inp      []string
		expected int
	}{
		{inp: []string{"apples", "apples"}, expected: 0},
		{inp: []string{}, expected: 0},
		{inp: []string{"apples", "flowers"}, expected: 2},
		{inp: []string{"apples"}, expected: 1},
	}
	for _, test := range tests {
		output := GetUniqueWords_optimized(test.inp)
		assert.Equal(t, test.expected, len(output))
	}
}
