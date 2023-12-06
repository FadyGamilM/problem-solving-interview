package tree

import (
	"errors"
)

type Node struct {
	Val    int
	Childs []*Node
}

// node-factory
func NewNode(val int) *Node {
	childs := make([]*Node, 0)
	return &Node{
		Val:    val,
		Childs: childs,
	}

}

// the receiver is not a pointer because i won't change anything in the node attributes
func (root Node) Sum_v1() int {
	queue := make([]Node, 0)
	queue = append(queue, root)
	total := 0
	frontNode := root
	for len(queue) > 0 {
		frontNode = queue[0]
		total += frontNode.Val
		for _, child := range frontNode.Childs {
			queue = append(queue, *child)
		}
		queue = queue[1:]
	}
	return total
}

// since in the question you didn't specify the type of the tree if its binary or a node might have multple childs, so started with the breadth first search which has a time complexity of O(n) and space complexity of O(W) where w is the maximum width of the tree and n is number of nodes in the tree
func SumBFS(node *Node) (int, error) {
	if node == nil {
		return -1, errors.New("cannot find sum to an empty tree")
	}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	total := 0
	frontNode := node
	for len(queue) > 0 {
		frontNode = queue[0]
		total += frontNode.Val
		queue = append(queue, frontNode.Childs...)
		queue = queue[1:]
	}
	return total, nil
}

// the complexity of DFS is :
// time = O()
// space = O(H) where H is the maximum height of the tree (Depth) and this space is on the call stack because i am following the recursive approach right here ..
func SumDFS(node *Node) (int, error) {
	// base case
	if node == nil {
		return -1, errors.New("cannot find sum to an empty tree/subtree")
	}

	total := node.Val

	// if this node has a child, we will do a DFS on each one of them , if not we will return the calculated sum so far to the previous call
	for _, child := range node.Childs {
		sumSoFar, err := SumDFS(child)
		// if we have an error due to the previous call [subtree is empty (node is nil)]
		if err != nil {
			// return 0 so the sum of this subtree = 0
			return 0, err
		}
		total += sumSoFar
	}
	return total, nil
}
