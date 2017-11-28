package ast

// Playlist declares a set of scenes that can be shown together.
type Playlist struct {
	Scenes          []PlaylistScene `json:"scenes"`
	DefaultDuration string          `json:"defaultDuration"`
}

// PlaylistScene declares a scene in a playlist, along with associated metadata.
type PlaylistScene struct {
	Scene    Scene  `json:"scene"`
	Duration string `json:"duration"`
}
