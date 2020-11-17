package player

type Option func(c *config)

type config struct {
	numLeds      uint
	maxFramerate int
}

var defaultConfig = config{
	numLeds:      50,
	maxFramerate: 60,
}

func WithNumLeds(numLeds uint) Option {
	return func(c *config) {
		c.numLeds = numLeds
	}
}

func WithMaxFramerate(maxFramerate int) Option {
	return func(c *config) {
		c.maxFramerate = maxFramerate
	}
}
