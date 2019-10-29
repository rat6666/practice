package linkedlist

type LinkedList struct {
	Length int
	First  *Node
	Last   *Node
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
