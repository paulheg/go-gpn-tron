package strategies

import gpntron "go-gpn-tron/internal/gpn-tron"

type Player struct {
	Name string

	// Last moves in order
	Tiles []Coordinate
}

func (p *Player) Direction() gpntron.Move {

	const windowSize = 5

	var size = len(p.Tiles)

	if size < windowSize && size >= 2 {
		return p.Tiles[size-1].Delta(p.Tiles[0]).Direction()
	} else if size >= windowSize {
		return p.Tiles[size-1].Delta(p.Tiles[size-5]).Direction()
	} else {
		return gpntron.Nothing
	}
}

func (p *Player) Position() Coordinate {
	return p.Tiles[len(p.Tiles)-1]
}

func isPlayerDangerous(self, enemy Player, chosen gpntron.Move) bool {

	return self.Position().Delta(enemy.Position()).Length() <= 2 && enemy.Direction().Opposite(chosen)

}
