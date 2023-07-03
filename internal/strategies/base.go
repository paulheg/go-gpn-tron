package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"log"
	"strings"
)

type Base struct {
	alive          bool
	height, width  int
	id             gpntron.PlayerID
	players        map[gpntron.PlayerID]*Player
	field          map[Coordinate]gpntron.PlayerID
	playerPosition Coordinate
}

func (b *Base) MessageOfTheDay(ex gpntron.Executor, message string) {
	log.Println("Message of the day:", message)
}

func (b *Base) Error(ex gpntron.Executor, err string) {
	log.Println("Error:", err)
}

func (b *Base) Game(ex gpntron.Executor, width int, height int, id gpntron.PlayerID) {
	b.width = width
	b.height = height
	b.id = id

	b.players = make(map[gpntron.PlayerID]*Player)
	b.field = make(map[Coordinate]gpntron.PlayerID)
}

func (b *Base) Position(ex gpntron.Executor, id gpntron.PlayerID, x int, y int) {
	b.field[Coordinate{x, y}] = id

	b.players[id].Tiles = append(b.players[id].Tiles, Coordinate{
		X: x,
		Y: y,
	})

	if id == b.id {
		b.playerPosition = Coordinate{
			X: x,
			Y: y,
		}
	}
}

func (b *Base) Player(ex gpntron.Executor, id gpntron.PlayerID, name string) {
	log.Println("Player", name, "joined")

	if _, ok := b.players[id]; !ok {
		b.players[id] = &Player{
			Name:  name,
			Tiles: make([]Coordinate, 0),
		}
	}
}

func (b *Base) Die(ex gpntron.Executor, ids ...gpntron.PlayerID) {
	for _, id := range ids {

		if id == b.id {
			b.alive = false
			log.Println("You died!")
		} else {
			log.Println(b.players[id].Name, "died")

			b.clearPlayer(id)
		}
	}
}

func (b *Base) Message(ex gpntron.Executor, id gpntron.PlayerID, message string) {
	log.Println("Message from", b.players[id], ":", message)
}

func (b *Base) Win(ex gpntron.Executor, wins int, losses int) {
	log.Println("You won!!!!", "Wins", wins, "Losses", losses)

	b.clear()
}

func (b *Base) Lose(ex gpntron.Executor, wins int, losses int) {
	log.Println("You lost :(", "Wins", wins, "Losses", losses)

	b.clear()
}

func (b *Base) clear() {
	for player := range b.players {
		delete(b.players, player)
	}

	for tile := range b.field {
		delete(b.field, tile)
	}

	b.alive = true
}

func (b *Base) clearPlayer(id gpntron.PlayerID) {

	for _, coord := range b.players[id].Tiles {
		delete(b.field, coord)
	}
}

func (b *Base) calcNextPos(move gpntron.Move) Coordinate {

	nextPos := Coordinate{
		X: b.playerPosition.X,
		Y: b.playerPosition.Y,
	}

	switch move {
	case gpntron.Down:
		nextPos.Y += 1
	case gpntron.Up:
		nextPos.Y -= 1
	case gpntron.Left:
		nextPos.X -= 1
	case gpntron.Right:
		nextPos.X += 1
	}

	if nextPos.X >= b.width {
		nextPos.X = 0
	}

	if nextPos.X < 0 {
		nextPos.X = b.width - 1
	}

	if nextPos.Y >= b.height {
		nextPos.Y = 0
	}

	if nextPos.Y < 0 {
		nextPos.Y = b.height - 1
	}

	return nextPos
}

func (b *Base) collision(move gpntron.Move) bool {
	_, ok := b.field[b.calcNextPos(move)]
	return ok
}

func (b *Base) PrintField() string {

	var buffer strings.Builder

	buffer.WriteRune('\n')

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {

			if v, ok := b.field[Coordinate{X: j, Y: i}]; ok {
				if v == b.id {
					buffer.WriteString("🟥")
				} else {
					buffer.WriteString("🟫")
				}
			} else {
				buffer.WriteString("⬛")
			}
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}
