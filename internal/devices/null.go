package devices

type Null struct {
}

func init() {
	Registry["null"] = &Null{}
}

func (n *Null) Setup(numLights int) {

}

func (n *Null) ShowColors(colors []uint32) {

}

func (n *Null) Clear() {

}

func (n *Null) Teardown() {

}
