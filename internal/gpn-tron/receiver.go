package gpntron

type Ticker interface {
	Tick(ex Executor) Move
}

type BaseReceiver interface {
	MessageOfTheDay(ex Executor, message string)
	Error(ex Executor, err string)
	Game(ex Executor, width, height int, id PlayerID)
	Position(ex Executor, id PlayerID, x, y int)
	Player(ex Executor, id PlayerID, name string)
	Die(ex Executor, ids ...PlayerID)
	Message(ex Executor, id PlayerID, message string)
	Win(ex Executor, wins int, losses int)
	Lose(ex Executor, wins int, losses int)
}

type Receiver interface {
	BaseReceiver
	Ticker
}
