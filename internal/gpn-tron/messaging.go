package gpntron

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const ParameterDelimeter = "|"
const PackageEnd = '\n'

func execute(input string, recv Receiver, ex Executor) {

	args := strings.Split(input, ParameterDelimeter)

	pType := PackageType(args[0])

	switch pType {
	case MessageOfTheDay:
		recv.MessageOfTheDay(ex, args[1])
	case Error:
		recv.Error(ex, args[1])
	case Game:
		width, height, id := 0, 0, 0

		width, _ = strconv.Atoi(args[1])
		height, _ = strconv.Atoi(args[2])
		id, _ = strconv.Atoi(args[3])

		recv.Game(ex, width, height, PlayerID(id))
	case Position:
		id, x, y := 0, 0, 0

		id, _ = strconv.Atoi(args[1])
		x, _ = strconv.Atoi(args[2])
		y, _ = strconv.Atoi(args[3])

		recv.Position(ex, PlayerID(id), x, y)
	case Player:
		id := 0
		name := args[2]

		id, _ = strconv.Atoi(args[1])

		recv.Player(ex, PlayerID(id), name)
	case Tick:
		move := recv.Tick(ex)
		ex.Move(move)
	case Die:
		ids := make([]PlayerID, len(args)-1)

		for i := 1; i < len(args); i++ {
			id, _ := strconv.Atoi(args[i])
			ids = append(ids, PlayerID(id))
		}

		recv.Die(ex, ids...)
	case Message:
		id := 0
		message := args[2]

		id, _ = strconv.Atoi(args[1])

		recv.Message(ex, PlayerID(id), message)
	case Win:
		wins, losses := 0, 0

		wins, _ = strconv.Atoi(args[1])
		losses, _ = strconv.Atoi(args[2])

		recv.Win(ex, wins, losses)
	case Lose:
		wins, losses := 0, 0

		wins, _ = strconv.Atoi(args[1])
		losses, _ = strconv.Atoi(args[2])

		recv.Lose(ex, wins, losses)
	default:
		log.Println("Unkown command", pType)
	}
}

func serialize(ptype PackageType, args ...interface{}) []byte {

	var buffer bytes.Buffer

	buffer.WriteString(string(ptype))
	if len(args) > 0 {
		buffer.WriteString(ParameterDelimeter)

		for i, arg := range args {
			buffer.WriteString(fmt.Sprint(arg))
			if i < len(args)-1 {
				buffer.WriteString(ParameterDelimeter)
			}
		}
	}

	buffer.WriteRune(PackageEnd)

	return buffer.Bytes()
}

type PackageType string

const (
	MessageOfTheDay PackageType = "motd"
	Join            PackageType = "join"
	Error           PackageType = "error"
	Game            PackageType = "game"
	Position        PackageType = "pos"
	Player          PackageType = "player"
	Tick            PackageType = "tick"
	Die             PackageType = "die"
	Moving          PackageType = "move"
	Chat            PackageType = "chat"
	Message         PackageType = "message"
	Win             PackageType = "win"
	Lose            PackageType = "lose"
)

type PlayerID int

type Move string

const (
	Up      Move = "up"
	Down    Move = "down"
	Left    Move = "left"
	Right   Move = "right"
	Nothing Move = "nothing"
)

func (mv Move) Opposite(m Move) bool {

	if mv != m {
		return (mv == Up && m == Down) ||
			(mv == Down && m == Up) ||
			(mv == Left && m == Right) ||
			(mv == Right && m == Left)
	}

	return false
}
