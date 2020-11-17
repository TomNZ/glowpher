package dsl

import (
	"time"

	"github.com/tomnz/glowpher/internal/variable"

	"github.com/tomnz/glowpher/internal/effect"
)

type Config struct {
	variables map[string]variable.Variable
	scenes    map[string]*Scene
	playlists map[string]*Playlist
}

type Playlist struct {
	name            string
	scenes          []PlaylistScene
	defaultDuration time.Duration
}

type PlaylistScene struct {
	name     string
	duration time.Duration
}

type Scene struct {
	name    string
	effects []effect.Effect
}
