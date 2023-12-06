package tree

type Tree struct {
	Root   *Node
	Length int
}

// tree factory
func NewTree(root *Node) *Tree {
	if root == nil {
		return &Tree{Length: 0}
	}
	return &Tree{
		Length: 1,
		Root:   root,
	}
}
