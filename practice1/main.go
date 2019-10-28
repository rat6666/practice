package main

import (
	"practice1/domain"
)

func main() {
	array := domain.CreateSlice(10)
	domain.PrintSlice(array)
	array = domain.FillDiag(array, 88)
	domain.PrintSlice(array)

	// domain.Drop(6, 1000)
}
