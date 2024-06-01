package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"log"
)

var _ gpntron.Receiver = (*FreeRealestate)(nil)

type FreeRealestate struct {
	Base
}

// Tick implements gpntron.Receiver.
func (f *FreeRealestate) Tick(ex gpntron.Executor) gpntron.Move {
	return f.HighestEmpty()
}

func (f *FreeRealestate) HighestEmpty() gpntron.Move {

	maxEmpty := uint(0)
	maxDir := gpntron.Nothing

	for _, direction := range gpntron.Directions {

		currentEmpty := uint(0)

		for i := 0; i < 5; i++ {
			nextCoord := f.playerPosition.Offset(direction, currentEmpty+1)

			playerId := f.Arena.Field(nextCoord)
			if playerId.Empty() {
				currentEmpty++
			} else {
				if currentEmpty > maxEmpty {
					maxEmpty = currentEmpty
					maxDir = direction
					break
				}
			}
		}
	}

	if maxDir == gpntron.Nothing {
		maxDir = f.ChooseNonColliding()
	}

	log.Printf("choosing %v", maxDir)
	return maxDir
}
