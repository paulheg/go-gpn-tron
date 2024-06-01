package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"math"
)

var (
	TopLeft     = []gpntron.Move{gpntron.Up, gpntron.Left}
	TopRight    = []gpntron.Move{gpntron.Up, gpntron.Right}
	BottomLeft  = []gpntron.Move{gpntron.Down, gpntron.Left}
	BottomRight = []gpntron.Move{gpntron.Down, gpntron.Right}

	Corners = [][]gpntron.Move{TopLeft, TopRight, BottomLeft, BottomRight}
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Equals(cord Coordinate) bool {
	return cord.X == c.X && cord.Y == c.Y
}

func (c Coordinate) OffsetD(directions ...gpntron.Move) Coordinate {
	for _, direction := range directions {
		c = c.Offset(direction, 1)
	}

	return c
}

func (c Coordinate) Offset(direction gpntron.Move, length uint) Coordinate {

	switch direction {
	case gpntron.Down:
		return c.Add(Coordinate{
			Y: int(length),
		})
	case gpntron.Left:
		return c.Add(Coordinate{
			X: -int(length),
		})
	case gpntron.Right:
		return c.Add(Coordinate{
			X: int(length),
		})
	case gpntron.Up:
		return c.Add(Coordinate{
			Y: -int(length),
		})
	default:
		return c
	}
}

func (c Coordinate) Add(cord Coordinate) Coordinate {
	return Coordinate{
		X: c.X + cord.X,
		Y: c.Y + cord.Y,
	}

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
