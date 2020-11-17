package player

import (
	"time"

	"github.com/tomnz/glowpher/internal/color"

	"github.com/tomnz/glowpher/internal/devices"
	"github.com/tomnz/glowpher/internal/dsl"
)

func Play(playlist *dsl.Playlist, scenes map[string]*dsl.Scene, device devices.Device, opts ...Option) {
	cfg := defaultConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	maxFrameDuration := time.Second / time.Duration(cfg.maxFramerate)

	scene := scenes[playlist.Scenes[0].Name]
	for {
		start := time.Now()
		colors := make([]color.Color, cfg.numLeds)
		for _, effect := range scene.Effects {

		}
	}
}
