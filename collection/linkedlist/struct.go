package linkedlist

//Collection LinkedList
type Collection struct {
	Length int
	First  *Element
	Last   *Element
}

//Element NODE
type Element struct {
	Value int
	Next  *Element
	Prev  *Element
}
