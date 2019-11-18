package game

import (
	"math/rand"
	"time"
)

// Player struct
type Player struct {
	Name  string
	Skill int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random gen
func (p Player) Random() bool {
	return p.Skill > (1 + rand.Intn(100))
}

// Play serve
func (p *Player) Play(chBall chan int, chWin chan string) {
	for {
		_, ok := <-chBall
		if ok {
			if p.Random() {
				chBall <- 1
			} else {
				close(chBall)
				break
			}
		} else {
			chWin <- p.Name
			break
		}
	}
}
