package config

// Scene declares an overall animated configuration.
type Scene struct {
	Name    string
	Effects []SceneEffect `json:"effects"`
}

// SceneEffect declares the type and parameters of a given effect instance in
// the scene.
type SceneEffect struct {
	Type   string           `json:"type"`
	Params map[string]Param `json:"params,omitempty"`
}
