package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
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
	r.lastMove = r.ChooseNonColliding()

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
