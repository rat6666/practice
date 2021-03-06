package linkedlist

import "fmt"

type LinkedList struct {
	Length int
	First  *Node
	Last   *Node
}

//AddNode Creating NODE
func (c *LinkedList) AddNode(value int) {
	newNode := &Node{}
	newNode.Value = value
	if c.Length == 0 {
		c.First = newNode
		c.Last = newNode
	}
	newNode.Prev = c.Last
	c.Last.Next = newNode
	c.Last = newNode
	c.Length++
}

func (c *LinkedList) Get(nodeIndex int) *Node {
	if nodeIndex >= c.Length {
		fmt.Printf("No element with index %d.\n", nodeIndex)
		return nil
	}
	nodeValue := c.First
	for i := 0; i < nodeIndex; i++ {
		nodeValue = nodeValue.Next
	}
	// fmt.Printf("The value of the element with index %d is %d.\n", nodeIndex, nodeValue.Value)
	return nodeValue
}

func (c *LinkedList) RemoveNode(removeIndex int) {
	if removeIndex > c.Length {
		fmt.Println("Index out of range")
		return
	}
	// remove := c.First
	// for i := 0; i < removeIndex; i++ {
	// 	remove = remove.Next
	// }
	remove := c.Get(removeIndex)
	remove.Next.Prev = remove.Prev
	remove.Prev.Next = remove.Next
	c.Length--
}

//PrintCollection Method for outputting NODE value
func (c *LinkedList) PrintCollection() {
	node := c.First
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
	fmt.Println()
}

func (c *LinkedList) FirstElement() *Node {
	return c.First
}

func (c *LinkedList) LastElement() *Node {
	return c.Last
}

func (c *LinkedList) LengthCollection() int {
	return c.Length
}
