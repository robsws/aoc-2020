package utils

// LinkedListNode - a node in a linked list
type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

// Insert - insert another linked list into this linked list at this point
func (node *LinkedListNode) Insert(ll *LinkedListNode) {
	oldNext := node.Next
	node.Next = ll
	end := ll.End()
	end.Next = oldNext
}

// Remove - remove n nodes from after the current node into their own separate LL
func (node *LinkedListNode) Remove(amount int) *LinkedListNode {
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
func (node *LinkedListNode) End() *LinkedListNode {
	endNode := node
	for endNode.Next != nil {
		endNode = endNode.Next
		if endNode == node {
			return node
		}
	}
	return endNode
}
