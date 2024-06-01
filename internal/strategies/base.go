package strategies

import (
	"cmp"
	gpntron "go-gpn-tron/internal/gpn-tron"
	"log"
	"slices"
)

var _ gpntron.BaseReceiver = (*Base)(nil)

type Base struct {
	Arena Arena

	alive          bool
	id             gpntron.PlayerID
	players        map[gpntron.PlayerID]*Player
	playerPosition Coordinate
}

func (b *Base) MessageOfTheDay(ex gpntron.Executor, message string) {
	log.Println("Message of the day:", message)
}

func (b *Base) Error(ex gpntron.Executor, err string) {
	log.Println("Error:", err)
}

func (b *Base) Game(ex gpntron.Executor, width int, height int, id gpntron.PlayerID) {
	b.id = id
	b.Arena = *NewArena(height, width)

	b.players = make(map[gpntron.PlayerID]*Player)
}

func (b *Base) Position(ex gpntron.Executor, id gpntron.PlayerID, x int, y int) {

	coord := Coordinate{X: x, Y: y}

	b.Arena.Set(coord, id)

	if id == b.id {
		b.playerPosition = coord
	}
}

func (b *Base) Player(ex gpntron.Executor, id gpntron.PlayerID, name string) {
	log.Println("Player", name, "joined")

	if _, ok := b.players[id]; !ok {
		b.players[id] = &Player{
			Name: name,
		}
	}
}

func (b *Base) Die(ex gpntron.Executor, ids ...gpntron.PlayerID) {
	for _, id := range ids {
		if id == b.id {
			b.alive = false
			log.Println("You died!")
		}
	}

	b.Arena.Clear(ids...)
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
		b.Arena.Clear(player)
		delete(b.players, player)
	}

	b.alive = true
}

func (b *Base) Collision(move gpntron.Move) bool {
	return b.CollisionLookahead(move, 1)
}

func (b *Base) CheckCorners() bool {
	for _, corner := range Corners {
		check := b.playerPosition.OffsetD(corner...)
		if !b.Arena.Field(check).Empty() {
			return true
		}
	}

	return false
}

func (b *Base) TunnelCollision(direction gpntron.Move, length uint) bool {
	moves := make([]gpntron.Move, length)

	for i := 0; i < int(length); i++ {
		moves[i] = direction
	}

	for i := 1; i < int(length); i++ {
		first := b.CollisionCheckMoves(append(moves[:i], direction.Tangential())...)
		second := b.CollisionCheckMoves(append(moves[:i], direction.Tangential().Opposite())...)

		if first && second {
			return true
		}
	}

	return false
}

func (b *Base) CollisionCheckMoves(moves ...gpntron.Move) bool {
	check := b.playerPosition.OffsetD(moves...)
	// log.Printf("checking %v", check)
	id := b.Arena.Field(check)

	return id.Empty()
}

func (b *Base) CollisionLookahead(move gpntron.Move, lookahead uint) bool {
	for distance := uint(1); distance <= lookahead; distance++ {
		check := b.playerPosition.Offset(move, distance)
		id := b.Arena.Field(check)
		// log.Printf("checking direction %s, my coord: %v with distance: %d, coordinate: %v, field value: %d", move, b.playerPosition, distance, check, id)

		if !id.Empty() {
			return true
		}
	}

	return false
}

func (b *Base) ChooseNonColliding() gpntron.Move {

	type scoredMove struct {
		Score int
		Move  gpntron.Move
	}

	possibleMoves := []scoredMove{}

	// select moves where we cant die
	for _, move := range gpntron.Directions {
		// skip if collision is up ahead
		if b.CollisionLookahead(move, 1) {
			continue
		} else {
			possibleMoves = append(possibleMoves, scoredMove{
				Score: 100,
				Move:  move,
			})
		}
	}

	for _, move := range possibleMoves {
		// check matching corners
		for _, corner := range Corners {
			if slices.Contains(corner, move.Move) {
				if b.CollisionCheckMoves(corner...) {
					move.Score /= 2
				}
			}
		}

		// reward paths with a lot of freedom
		for i := 10; i >= 2; i-- {
			if !b.CollisionLookahead(move.Move, uint(i)) {
				move.Score += i * 4
				break
			}
		}

		// check tunneling
	}

	if len(possibleMoves) == 0 {
		log.Print("were dead :(")
		return gpntron.Nothing
	}

	best := slices.MaxFunc[[]scoredMove](possibleMoves, func(a, b scoredMove) int {
		return cmp.Compare(a.Score, b.Score)
	})
	log.Printf("chosing %s with score %d", best.Move, best.Score)

	return best.Move
}
