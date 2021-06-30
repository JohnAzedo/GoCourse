package tree

/*
Node is struct to binary tree
@attrs:
	- Number int: Data in node
	- Right *Node: Pointer to tree right side
	- Left *Node: Pointer to tree left side
*/
type Node struct {
	Number int
	Right *Node 
	Left *Node
}