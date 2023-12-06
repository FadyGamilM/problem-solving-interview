package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		inp            string
		expectedOutput string
	}{
		{inp: "åbba", expectedOutput: "å1b2a1"}, // test with wierd characters
		{inp: "", expectedOutput: ""},           // test with wierd characters
		{inp: "x", expectedOutput: "x1"},        // test with wierd characters
		{inp: "5", expectedOutput: "51"},        // test with wierd characters
	}
	sw1 := NewSlidingWindow()
	for _, test := range tests {
		output := sw1.RunLengthEncode(test.inp)
		assert.Equal(t, test.expectedOutput, output)
		sw1.Reset()
	}
}
