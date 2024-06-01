package visualize

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"log"
	"strings"
	"time"
)

type Emotes string

const (
	Panda  Emotes = "🐼"
	Turtle Emotes = "🐢"
	Yellow Emotes = "🟨"
	Green  Emotes = "🟩"
	Blue   Emotes = "🟦"
	Purple Emotes = "🟪"
	Brown  Emotes = "🟫"
	White  Emotes = "⬜"
)

var directionArrows = map[gpntron.Move]string{
	gpntron.Down:    "⬇️",
	gpntron.Up:      "⬆️",
	gpntron.Left:    "⬅️",
	gpntron.Right:   "➡️",
	gpntron.Nothing: "🔃",
}

var Blocks = []Emotes{
	Yellow, Green, Blue, Purple,
}

var Numbers = []Emotes{
	Emotes("🅰️"), Emotes("1️⃣"), Emotes("2️⃣"), Emotes("3️⃣"), Emotes("4️⃣"),
	Emotes("5️⃣"), Emotes("6️⃣"), Emotes("7️⃣"), Emotes("8️⃣"), Emotes("9️⃣"),
	Emotes("🔟"),
}

type consoleVisualizer struct {
	gpntron.BaseReceiver
	gpntron.Ticker
	*strategies.Arena

	myID gpntron.PlayerID
}

func NewConsoleVisualizer(base gpntron.BaseReceiver, ticker gpntron.Ticker, arena *strategies.Arena) Visualizer {
	return &consoleVisualizer{
		BaseReceiver: base,
		Ticker:       ticker,
		Arena:        arena,
	}
}

func (v *consoleVisualizer) Die(ex gpntron.Executor, ids ...gpntron.PlayerID) {

	log.Printf("players died: %v", ids)

	v.BaseReceiver.Die(ex, ids...)
}

func (v *consoleVisualizer) Tick(ex gpntron.Executor) gpntron.Move {
	begin := time.Now()
	log.Println(v.Print())

	defer log.Printf("Calculation Time: %d ms", time.Now().Sub(begin).Milliseconds())
	return v.Ticker.Tick(ex)
}

func (v *consoleVisualizer) Game(ex gpntron.Executor, width int, height int, id gpntron.PlayerID) {
	v.myID = id
	log.Printf("Im Player %d", v.myID)

	v.BaseReceiver.Game(ex, width, height, id)
}

func (b *consoleVisualizer) Print() string {

	var buffer strings.Builder

	buffer.WriteRune('\n')

	for i := 0; i < b.Arena.Size().X; i++ {
		for j := 0; j < b.Arena.Size().Y; j++ {

			coord := strategies.Coordinate{X: j, Y: i}

			field := b.Arena.Field(coord)
			if field.Empty() {
				buffer.WriteString("⬛")
			} else {
				// draw an arrow at the end
				if playerId := b.Arena.IsHead(coord); !playerId.Empty() {
					buffer.WriteString(directionArrows[b.Arena.Moves()[playerId].Direction()])
				} else {
					block := '🍇' + field
					if field == b.myID {
						block = '🐧'
					}

					buffer.WriteString(string(rune(block)))
				}
			}
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}
