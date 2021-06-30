package ll

// Regular linked list

import (
	"fmt"
	"strings"
)

// Node to double linked list
type Node struct {
	Data     int64 // Number
	Next     *Node // Next Node
}

// CreateList creates a list of nodes with values ​​passed by parameters
func CreateList(values ...int64) *Node {
	if len(values) > 0 {
		head := Node{values[0], nil}
		aux := &head
		// O(n)
		for i, value := range values {
			if i > 0 {
				tail := Node{value, nil}
				aux.Next = &tail
				aux = &tail
			}
		}
		return &head
	}
	return nil
}

// CountNodes counts the number of nodes in the linked list
func (head Node) CountNodes() int{
	ptr := &head
	counter := 0

	// O(n)
	for ptr != nil {
		ptr = ptr.Next
		counter++
	}

	return counter
}

// PrintList list all nodes
func (head Node) PrintList() string {
	var out string = ""
	ptr := &head
	// O(n)
	for ptr != nil {
		out = fmt.Sprintf("%v(%v, %p)->", out, ptr.Data, ptr.Next)
		ptr = ptr.Next
	}

	out = strings.ReplaceAll(out, "0x0", "<nil>")
	out = out + "<nil>"
	return out
}

// Append add a new node at end of list
func (head Node) Append(number int64) {
	//TODO -> Allow the user to choose the location of the next node (#1)
	// If using Append function in CreateList, the complexity will be O(nˆ2)
	tail := Node{number, nil}
	ptr := &head

	// Last node O(n)
	for ptr.Next != nil {
		ptr = ptr.Next
	}

	ptr.Next = &tail
}

// ChangeHead add new node at the beginning of list
func (head *Node) ChangeHead(number int64){
	data := head.Data
	ptr := head.Next
	head.Data = number
	head.Next = &Node{data, ptr}
}

// RemoveNode Removes any node from list
func (head *Node) RemoveNode(number int64) {
	// Redirect the selected node to the end of list and remove it
	ptr := head

	for ptr != nil {
		if ptr.Data == number{
			ptr.Data = ptr.Next.Data
			number = ptr.Data
			// Check if next node is the last and delete it from linked list 
			if ptr.Next.Next == nil {
				ptr.Next = nil
			}
		}
		ptr = ptr.Next	
	}
}