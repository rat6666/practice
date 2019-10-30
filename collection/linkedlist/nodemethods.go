package linkedlist

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
