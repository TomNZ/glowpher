package config

// Playlist declares a set of scenes that can be shown together.
type Playlist struct {
	Scenes          []Scene `json:"scenes"`
	DefaultDuration string  `json:"defaultDuration"`
}
