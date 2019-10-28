package domain

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Drop(x int, y int) {
	slice := make(map[int]float32)
	z := 0
	z1 := 0
	for i := 0; i < y; i++ {
		z = rand.Intn(x)
		z1 = rand.Intn(x)
		r := z + z1 + 2
		slice[r]++
	}
	for index, value := range slice {
		value = (value / float32(y)) * 100
		fmt.Println("number ", index, "-", value, " %")
	}
}
