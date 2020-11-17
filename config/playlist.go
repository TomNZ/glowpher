package config

// Playlist declares a set of scenes that can be shown together.
type Playlist struct {
	Name            string          `json:"name"`
	Scenes          []PlaylistScene `json:"scenes"`
	DefaultDuration string          `json:"defaultDuration"`
}

// PlaylistScene declares a scene within a dsl.
type PlaylistScene struct {
	Name     string  `json:"name"`
	Duration *string `json:"duration,omitempty"`
}
