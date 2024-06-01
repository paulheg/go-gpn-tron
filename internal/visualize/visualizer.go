package visualize

import gpntron "go-gpn-tron/internal/gpn-tron"

type Visualizer interface {
	gpntron.Receiver
}
