package main

import "fmt"

//Problem 09
//Given the below tree structure, write a function sum that accepts a node and returns the
//sum of integers for the whole tree represented by the given node argument
//
//struct Node {
//value: Integer,
//children: [Node] # array of nodes
//}

type Node struct {
	Value    int
	Children []*Node
}

func calculateSum(node *Node) int {
	if node == nil {
		return 0
	}

	sum := node.Value

	for _, child := range node.Children {
		sum += calculateSum(child)
	}

	return sum
}

func main() {

	node7 := &Node{Value: 7, Children: []*Node{}}
	node8 := &Node{Value: 8, Children: []*Node{}}
	node9 := &Node{Value: 9, Children: []*Node{}}
	node5 := &Node{Value: 5, Children: []*Node{}}
	node6 := &Node{Value: 6, Children: []*Node{}}
	node4 := &Node{Value: 4, Children: []*Node{}}
	node3 := &Node{Value: 3, Children: []*Node{node7, node8, node9}}
	node2 := &Node{Value: 2, Children: []*Node{node4, node5, node6}}
	node1 := &Node{Value: 1, Children: []*Node{node2, node3}}

	fmt.Println("Sum of the tree:", calculateSum(node1)) // Output: 45 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9)
}
