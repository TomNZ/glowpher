package devices

import "github.com/tomnz/glowpher/internal/color"

type Null struct {
}

func init() {
	Registry["null"] = &Null{}
}

func (n *Null) Setup(numLights int) error {
	return nil
}

func (n *Null) ShowColors(colors []color.Color) error {
	return nil
}

func (n *Null) Clear() error {
	return nil
}

func (n *Null) Teardown() {
}
