package linkedlist

// Node - a node in a linked list
type Node struct {
	Value int
	Next  *Node
}

// Insert - insert another linked list into this linked list at this point
func (node *Node) Insert(ll *Node) {
	oldNext := node.Next
	node.Next = ll
	end := ll.End()
	end.Next = oldNext
}

// Remove - remove n nodes from after the current node into their own separate LL
func (node *Node) Remove(amount int) *Node {
	nodePtr := node
	for i := 0; i < amount; i++ {
		nodePtr = nodePtr.Next
	}
	newNext := nodePtr.Next
	nodePtr.Next = nil
	oldNext := node.Next
	node.Next = newNext
	return oldNext
}

// End - navigate to the end of the linked list
func (node *Node) End() *Node {
	endNode := node
	for endNode.Next != nil {
		endNode = endNode.Next
		if endNode == node {
			return node
		}
	}
	return endNode
}
