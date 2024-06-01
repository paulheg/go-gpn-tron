package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"sync"
)

type Moves []Coordinate

func (m Moves) Head() Coordinate {
	return m[len(m)-1]
}

func (p Moves) Direction() gpntron.Move {

	const windowSize = 5

	size := len(p)

	if size >= 2 {
		return p[size-1].Delta(p[size-2]).Direction()
	}

	return gpntron.Nothing
}

func NewArena(height, width int) *Arena {
	a := &Arena{
		height:      height,
		width:       width,
		field:       make([]gpntron.PlayerID, height*width),
		playerMoves: make(map[gpntron.PlayerID]Moves),
	}

	for i := 0; i < len(a.field); i++ {
		a.field[i] = -1
	}

	return a
}

type Arena struct {
	sync.Mutex

	height, width int
	field         []gpntron.PlayerID
	playerMoves   map[gpntron.PlayerID]Moves
}

func (b *Arena) Moves() map[gpntron.PlayerID]Moves {
	return b.playerMoves
}

func (b *Arena) Field(coord Coordinate) gpntron.PlayerID {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	coord = b.WrapCoord(coord)
	return b.field[b.convertCoord(coord)]
}

func (b *Arena) Set(coord Coordinate, player gpntron.PlayerID) {
	b.Mutex.Lock()
	coord = b.WrapCoord(coord)

	b.field[b.convertCoord(coord)] = player
	b.playerMoves[player] = append(b.playerMoves[player], coord)
	b.Mutex.Unlock()
}

// Wrap the coordinate around the map if out of bounds
func (b *Arena) WrapCoord(c Coordinate) Coordinate {

	if b.width > 0 {
		c.X = (c.X + 2*b.width) % b.width
	}

	if b.height > 0 {
		c.Y = (c.Y + 2*b.height) % b.height
	}

	return c
}

func (a *Arena) convertCoord(c Coordinate) int {
	return c.Y*a.width + c.X
}

func (a *Arena) Clear(players ...gpntron.PlayerID) {
	a.Lock()

	for _, player := range players {
		for _, coord := range a.playerMoves[player] {
			a.field[a.convertCoord(coord)] = -1
		}

		delete(a.playerMoves, player)
	}

	a.Unlock()
}

func (a *Arena) Size() Coordinate {
	return Coordinate{X: a.width, Y: a.height}
}

func (a *Arena) IsHead(c Coordinate) gpntron.PlayerID {
	for playerId, moves := range a.playerMoves {
		if moves.Head().Equals(c) {
			return playerId
		}
	}

	return -1
}
