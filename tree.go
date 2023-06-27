package main

import (
	"fmt"
)

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

// Function to recursively print the tree
func printTree(node *TreeNode, indent string) {
	fmt.Println(indent, node.Value)

	// Recursively print the children nodes
	for _, child := range node.Children {
		printTree(child, indent+"  ")
	}
}
func main() {
	// Create a sample tree
	root := &TreeNode{Value: 1}
	node2 := &TreeNode{Value: 2}
	node3 := &TreeNode{Value: 3}
	node4 := &TreeNode{Value: 4}
	node5 := &TreeNode{Value: 5}

	root.Children = []*TreeNode{node2, node3}
	node2.Children = []*TreeNode{node4}
	node3.Children = []*TreeNode{node5}

	// Print the tree
	printTree(root, "")
}
