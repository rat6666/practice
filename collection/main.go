package main

import (
	"collection/linkedlist"
	"fmt"
)

func main() {
	collection := linkedlist.LinkedList{}
	for i := 2; i < 10; i++ {
		collection.AddNode(i)
	}
	collection.PrintCollection()
	collection.AddNode(55)
	collection.PrintCollection()
	fmt.Println(collection.Get(3).GetValue())
	fmt.Println(collection.Get(3).NextNode().GetValue())
	fmt.Println(collection.Get(3).PrevNode().GetValue())
	collection.Get(15)
	collection.FirstElement()
	collection.LastElement()
	collection.LengthCollection()
	collection.RemoveNode(3)
	collection.PrintCollection()
	collection.LengthCollection()
	collection.RemoveNode(55)

}
