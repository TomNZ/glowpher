package playlist

import (
	"time"

	"github.com/tomnz/glowpher/internal/effect"
)

type Playlist struct {
	scenes          []*Scene
	defaultDuration time.Duration
}

type Scene struct {
	name     string
	effects  []effect.Effect
	duration time.Duration
}
