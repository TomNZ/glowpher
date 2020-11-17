package dsl

import (
	"time"

	"github.com/tomnz/glowpher/internal/variable"

	"github.com/tomnz/glowpher/internal/effect"
)

type Config struct {
	Variables map[string]variable.Variable
	Scenes    map[string]*Scene
	Playlists map[string]*Playlist
}

type Playlist struct {
	Name            string
	Scenes          []PlaylistScene
	DefaultDuration time.Duration
}

type PlaylistScene struct {
	Name     string
	Duration time.Duration
}

type Scene struct {
	Name    string
	Effects []effect.Effect
}
