package main

import (
	"fmt"
	"tennis/game"
)

func main() {
	p1 := game.Player{Name: "Vasya", Skill: 50}
	p2 := game.Player{Name: "Petya", Skill: 50}
	match := game.Match{P1: p1, P2: p2}
	for i := 0; i < 10; i++ {
		fmt.Println(match.Start())
	}
}
