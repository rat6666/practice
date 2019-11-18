package game

import (
	"fmt"
)

// Match struct
type Match struct {
	P1, P2 Player
}

// Start (one serve)
func (m *Match) Start() string {
	chBall := make(chan int)
	chWin := make(chan string)

	go m.P1.Play(chBall, chWin)

	go m.P2.Play(chBall, chWin)

	chBall <- 1

	winner, _ := <-chWin

	return fmt.Sprintf("Player %v win", winner)
}
