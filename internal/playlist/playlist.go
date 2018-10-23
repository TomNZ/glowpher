package playlist

import (
	"github.com/tomnz/glowpher/internal/variable"
	"time"

	"github.com/tomnz/glowpher/internal/effect"
)

type Playlist struct {
	scenes          []*Scene
	variables       map[string]variable.Variable
	defaultDuration time.Duration
}

type Scene struct {
	name     string
	effects  []effect.Effect
	duration time.Duration
}
