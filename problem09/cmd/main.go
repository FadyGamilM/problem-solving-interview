package main

import (
	"halanproblem09/tree"
	"log"
)

func main() {
	rootNode := tree.NewNode(1)
	node2 := tree.NewNode(2)
	node3 := tree.NewNode(3)
	node4 := tree.NewNode(4)
	rootNode.Childs = []*tree.Node{node2, node3, node4}
	node5 := tree.NewNode(5)
	node2.Childs = []*tree.Node{node5}
	node6 := tree.NewNode(6)
	node3.Childs = []*tree.Node{node6}
	node7 := tree.NewNode(7)
	node8 := tree.NewNode(8)
	node4.Childs = []*tree.Node{node7, node8}
	testTree := tree.NewTree(rootNode)
	log.Println(tree.SumBFS(testTree.Root))
	log.Println(tree.SumDFS(testTree.Root))
}
