package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSum(t *testing.T) {
	rootNode := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	rootNode.Childs = []*Node{node2, node3, node4}
	node5 := NewNode(5)
	node2.Childs = []*Node{node5}
	node6 := NewNode(6)
	node3.Childs = []*Node{node6}
	node7 := NewNode(7)
	node8 := NewNode(8)
	node4.Childs = []*Node{node7, node8}
	tree := NewTree(rootNode)
	testCases := []struct {
		root        *Node
		expectedSum int
	}{
		{
			root:        tree.Root,
			expectedSum: 36,
		},
	}

	for _, test := range testCases {
		sum, err := SumBFS(tree.Root)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedSum, sum)

		dfsSum, err := SumDFS(tree.Root)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedSum, dfsSum)
	}
}

func TestTreeSumWithOneNode(t *testing.T) {
	rootNode := NewNode(50)
	tree := NewTree(rootNode)
	testCases := []struct {
		root        *Node
		expectedSum int
	}{
		{
			root:        tree.Root,
			expectedSum: 50,
		},
	}
	for _, test := range testCases {
		sum, err := SumBFS(tree.Root)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedSum, sum)

		dfsSum, err := SumDFS(tree.Root)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedSum, dfsSum)
	}
}

func TestTreeSumWithEmptyTree(t *testing.T) {
	tree := NewTree(nil)
	testCases := []struct {
		root        *Node
		expectedSum int
	}{
		{
			root:        tree.Root,
			expectedSum: -1,
		},
	}
	for _, test := range testCases {
		sum, err := SumBFS(tree.Root)
		assert.Error(t, err)
		assert.Equal(t, test.expectedSum, sum)

		dfsSum, err := SumDFS(tree.Root)
		assert.Error(t, err)
		assert.Equal(t, test.expectedSum, dfsSum)
	}
}
