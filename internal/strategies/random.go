package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"log"
	"math/rand"
)

var _ gpntron.Receiver = &Random{}

type Random struct {
	Base
	n        int
	lastMove gpntron.Move
}

// Tick implements gpntron.Receiver.
func (r *Random) Tick(ex gpntron.Executor) gpntron.Move {
	if len(r.lastMove) == 0 || r.lastMove == gpntron.Nothing {
		r.lastMove = r.randomMove()
	}

	if r.n > 3 {
		r.lastMove = r.randomMove()
		r.n = 0
	}

	for id, player := range r.players {
		if id != r.id {
			if isPlayerDangerous(*r.players[r.id], *player, r.lastMove) {
				r.lastMove = r.randomMove()
			}
		}
	}

	for r.collision(r.lastMove) {
		r.lastMove = r.randomMove()
		r.n = 0
	}

	r.n += 1

	log.Println(r.PrintField())
	return r.lastMove
}

func (r *Random) randomMove() gpntron.Move {

	switch rand.Intn(4) {
	case 0:
		return gpntron.Down
	case 1:
		return gpntron.Left
	case 2:
		return gpntron.Right
	case 3:
		return gpntron.Up
	default:
		return gpntron.Nothing
	}
}
