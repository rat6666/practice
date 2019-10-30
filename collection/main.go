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
	fmt.Println("Output created collection")
	collection.PrintCollection()
	fmt.Println("Adding new element to collection")
	collection.AddNode(55)
	fmt.Println("Output changed collection")
	collection.PrintCollection()
	fmt.Println("Getting element value through index")
	fmt.Println(collection.Get(3).GetValue())
	fmt.Println("Output next element")
	fmt.Println(collection.Get(3).NextNode().GetValue())
	fmt.Println("Output previous element")
	fmt.Println(collection.Get(3).PrevNode().GetValue())
	fmt.Println("Trying to get the value of an element with an index out of range")
	collection.Get(15)
	collection.FirstElement()
	collection.LastElement()
	collection.LengthCollection()
	fmt.Println("Deleting element through index")
	collection.RemoveNode(3)
	collection.PrintCollection()
	fmt.Println("Trying to delete an element with an index out of range")
	collection.RemoveNode(55)

}
