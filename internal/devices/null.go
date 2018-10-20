package devices

type Null struct {
}

func init() {
	Registry["null"] = &Null{}
}

func (n *Null) Setup(numLights int) error {
	return nil
}

func (n *Null) ShowColors(colors []uint32) error {
	return nil
}

func (n *Null) Clear() error {
	return nil
}

func (n *Null) Teardown() {
}
