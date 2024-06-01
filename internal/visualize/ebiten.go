package visualize

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 500
	screenHeight = 500
)

var enemyColors = []color.Color{
	color.RGBA{R: 0x24, G: 0x7B, B: 0x7B, A: 255}, // #247B7B
	color.RGBA{R: 0x78, G: 0xCD, B: 0xD7, A: 255}, // #78CDD7
	color.RGBA{R: 0x44, G: 0xA1, B: 0xA0, A: 255}, // #44A1A0
	color.RGBA{R: 0x0D, G: 0x5C, B: 0x63, A: 255}, // #0D5C63
	color.RGBA{R: 0x0D, G: 0xFF, B: 0x63, A: 255}, // #0D5C63
	color.RGBA{R: 0x81, G: 0xF0, B: 0xE5, A: 255}, // #81F0E5
	color.RGBA{R: 0x73, G: 0x19, B: 0x63, A: 255}, // #731963
}

type EbitenVisualizer interface {
	ebiten.Game
	Visualizer
	Close()
}

func NewEbitenVisualizer(ticker gpntron.Ticker) EbitenVisualizer {
	return &ebitenVisualizer{
		Ticker: ticker,
		done:   false,
	}
}

type ebitenVisualizer struct {
	strategies.Base
	gpntron.Ticker
	done bool

	lastUpdate        time.Time
	waitingCircleSize int
}

// Close implements EbitenVisualizer.
func (e *ebitenVisualizer) Close() {
	e.done = true
}

// Layout implements EbitenVisualizer.
func (e *ebitenVisualizer) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update implements EbitenVisualizer.
func (e *ebitenVisualizer) Update() error {

	delta := e.lastUpdate.Sub(time.Now())

	if e.done {
		return ebiten.Termination
	}

	e.waitingCircleSize = ((e.waitingCircleSize + 1*int(delta.Milliseconds())) % 200)

	e.lastUpdate = time.Now()
	return nil
}

// Draw implements EbitenVisualizer.
func (e *ebitenVisualizer) Draw(screen *ebiten.Image) {

	vector.DrawFilledRect(screen, 0, 0, screenWidth, screenHeight, color.White, false)
	vector.StrokeRect(screen, 0, 0, screenHeight, screenWidth, .2, enemyColors[0], false)

	size := e.Arena.Size()

	// draw loading screen
	if size.Y == 0 || size.X == 0 {
		vector.DrawFilledCircle(screen,
			screenWidth/2, screenHeight/2, float32(e.waitingCircleSize), enemyColors[0], false)

		return
	}

	// draw actual game
	blockHeight, blockWidth := screen.Bounds().Dy()/size.Y, screen.Bounds().Dx()/size.X

	for id, moves := range e.Arena.Moves() {
		for _, move := range moves {
			color := enemyColors[int(id)%len(enemyColors)]

			vector.DrawFilledRect(screen, float32(move.X*blockWidth), float32(move.Y*blockHeight), float32(blockWidth), float32(blockHeight), color, false)
		}

	}
}
