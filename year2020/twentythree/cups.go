package twentythree

import (
	"aoc-go/files"
	"aoc-go/utils"
	"strconv"
	"strings"
)

type cupCircleLL struct {
	currentNode  *utils.LinkedListNode
	nodePointers map[int]*utils.LinkedListNode
}

func parseCups(filename string, fillto int) cupCircleLL {
	/* Create a circular linked list and a map of every
	   value to the address of its linked list node */
	lines := files.GetLines(filename)
	cupsStr := strings.Split(lines[0], "")
	cupsInt := make([]int, len(cupsStr))
	for i, cup := range cupsStr {
		cupsInt[i] = utils.MustAtoi(cup)
	}
	nodePointers := make(map[int]*utils.LinkedListNode)
	currentNode := utils.LinkedListNode{Value: cupsInt[0], Next: nil}
	firstNodePtr := &currentNode
	nodePointers[cupsInt[0]] = &currentNode
	nodePtr := &currentNode
	for i := 1; i < len(cupsInt); i++ {
		nextNode := utils.LinkedListNode{Value: cupsInt[i], Next: nil}
		nodePtr.Next = &nextNode
		nodePointers[cupsInt[i]] = &nextNode
		nodePtr = &nextNode
	}
	// If fillto is set, fill up the circle with sequential numbers
	if fillto >= len(cupsInt) {
		for i := len(cupsInt) + 1; i <= fillto; i++ {
			nextNode := utils.LinkedListNode{Value: i, Next: nil}
			nodePtr.Next = &nextNode
			nodePointers[i] = &nextNode
			nodePtr = &nextNode
		}
	}
	nodePtr.Next = firstNodePtr
	return cupCircleLL{firstNodePtr, nodePointers}
}

func (c *cupCircleLL) doRound() {
	// Remove next three cups
	subList := c.currentNode.Remove(3)
	// Calculate the destination value for insertion of those cups
	destValue := c.currentNode.Value - 1
	if destValue < 1 {
		// Once destination falls below 0, reset to highest value
		destValue = len(c.nodePointers)
	}
	gotDestValue := false
	for !gotDestValue {
		gotDestValue = true
		node := subList
		for node != nil {
			if node.Value == destValue {
				destValue = destValue - 1
				if destValue < 1 {
					// Once destination falls below 0, reset to highest value
					destValue = len(c.nodePointers)
				}
				gotDestValue = false
				break
			}
			node = node.Next
		}
	}
	// Find the node that has that destination value
	dest := c.nodePointers[destValue]
	// Insert the cups into the new destination
	dest.Insert(subList)
	// Move the current cup on one
	c.currentNode = c.currentNode.Next
}

func (c cupCircleLL) serialize() string {
	str := ""
	node := c.nodePointers[1].Next
	for node.Value != 1 {
		str += strconv.Itoa(node.Value)
		node = node.Next
	}
	return str
}
