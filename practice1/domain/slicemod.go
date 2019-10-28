package domain

import (
	"fmt"
)

func CreateSlice(y int) [][]int {
	newSlice := make([][]int, y)
	for i := 0; i < y; i++ {
		newSlice[i] = make([]int, y)
	}
	return newSlice
}

func PrintSlice(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println()
}

func FillDiag(s [][]int, fill int) [][]int {
	for i := 0; i < len(s); i++ {
		s[i][i] = fill
		s[i][len(s)-i-1] = fill
	}
	return s
}
