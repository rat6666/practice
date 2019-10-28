package linkedlist

import "fmt"

//AddElement Creating NODE
func (c *Collection) AddElement(value int) {
	newElement := &Element{}
	newElement.Value = value
	if c.Length == 0 {
		c.First = newElement
		c.Last = newElement
		c.Length++
		return
	}
	c.Last.Next = newElement
	c.Last = newElement
	c.Length++
}

//PrintCollection Method for outputting NODE value
func (c *Collection) PrintCollection() {
	result := c.First
	for result != nil {
		fmt.Print(result.Value, " ")
		result = result.Next
	}
	fmt.Println()
}
