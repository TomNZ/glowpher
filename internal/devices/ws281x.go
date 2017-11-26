// +build rpiws281x

package devices

import "github.com/tomnz/glowpher/internal/rpiws281x"

func init() {
	Registry["ws281x"] = &rpiws281x.WS281x{}
}
