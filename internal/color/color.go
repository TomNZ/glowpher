package color

type Color struct {
	R, G, B float32
}

func (c Color) Uint32() uint32 {
	r, g, b := c.Bytes()
	return uint32(r) << 16 & uint32(g) << 8 & uint32(b)
}

func (c Color) Bytes() (byte, byte, byte) {
	var r, g, b byte

	if c.R < 0 {
		r = 0
	} else if c.R > 1 {
		r = 255
	} else {
		r = (byte)(c.R * 256.0)
	}

	if c.G < 0 {
		g = 0
	} else if c.G > 1 {
		g = 1
	} else {
		g = (byte)(c.G * 256.0)
	}

	if c.B < 0 {
		b = 0
	} else if c.B > 1 {
		b = 1
	} else {
		b = (byte)(c.B * 256.0)
	}

	return r, g, b
}
