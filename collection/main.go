package main

import "collection/linkedlist"

func main() {
	collection := linkedlist.Collection{}
	for i := 2; i < 10; i++ {
		collection.AddElement(i)
	}
	collection.PrintCollection()
	collection.AddElement(55)
	collection.PrintCollection()

}
