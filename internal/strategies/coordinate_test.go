package strategies_test

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinateOffset(t *testing.T) {
	testCases := []struct {
		desc      string
		a         strategies.Coordinate
		direction gpntron.Move
		expected  strategies.Coordinate
		offset    uint
	}{
		{
			desc: "up",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			direction: gpntron.Up,
			expected: strategies.Coordinate{
				X: 1,
				Y: 0,
			},
			offset: 1,
		},
		{
			desc: "down",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			direction: gpntron.Down,
			expected: strategies.Coordinate{
				X: 1,
				Y: 2,
			},
			offset: 1,
		},
		{
			desc: "left",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			direction: gpntron.Left,
			expected: strategies.Coordinate{
				X: 0,
				Y: 1,
			},
			offset: 1,
		},
		{
			desc: "right",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			direction: gpntron.Right,
			expected: strategies.Coordinate{
				X: 2,
				Y: 1,
			},
			offset: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.a.Offset(tC.direction, tC.offset)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestCoordinateAdd(t *testing.T) {
	testCases := []struct {
		desc     string
		a, b     strategies.Coordinate
		expected strategies.Coordinate
	}{
		{
			desc: "basic add",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			b: strategies.Coordinate{
				X: 2,
				Y: 2,
			},
			expected: strategies.Coordinate{
				X: 3,
				Y: 3,
			},
		},
		{
			desc: "",
			a: strategies.Coordinate{
				X: 1,
				Y: 1,
			},
			b: strategies.Coordinate{
				X: -5,
				Y: -5,
			},
			expected: strategies.Coordinate{
				X: -4,
				Y: -4,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.a.Add(tC.b)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
