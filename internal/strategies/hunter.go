package strategies

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
)

var _ gpntron.Receiver = &Hunter{}

type Hunter struct {
	Base
}

func (h *Hunter) Tick(ex gpntron.Executor) gpntron.Move {
	return gpntron.Down
}
