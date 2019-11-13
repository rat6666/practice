package linkedlist

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

//GetValue Return NODE value
func (e *Node) GetValue() int {
	return e.Value
}
func (e *Node) NextNode() *Node {
	return e.Next
}
func (e *Node) PrevNode() *Node {
	return e.Prev
}
