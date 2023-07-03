package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"math"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Delta(coord Coordinate) Coordinate {
	return Coordinate{
		X: c.X - coord.X,
		Y: c.Y - coord.Y,
	}
}

func (c Coordinate) Length() int {
	absX := int(math.Abs(float64(c.X)))
	absY := int(math.Abs(float64(c.Y)))

	if absX > absY {
		return absX
	}

	return absY
}

func (c Coordinate) Direction() gpntron.Move {

	absX := math.Abs(float64(c.X))
	absY := math.Abs(float64(c.Y))

	if absX > absY {
		// left or right

		if c.X > 0 {
			return gpntron.Right
		} else {
			return gpntron.Left
		}

	} else {
		// up or down
		if c.Y > 0 {
			return gpntron.Down
		} else {
			return gpntron.Up
		}

	}
}
