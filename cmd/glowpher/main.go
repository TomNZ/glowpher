package main

import (
	"encoding/json"
	"github.com/k0kubun/pp"
	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/devices"
	"github.com/tomnz/glowpher/internal/playlist"
	"io/ioutil"
	"log"
	"math/rand"
)

const pixels = 52

func main() {
	var cfg config.Config

	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(configFile, &cfg); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}

	pl, err := playlist.Compile(cfg)
	if err != nil {
		log.Fatalf("couldn't compile playlist: %s", err)
	}

	pp.Print(pl)

	dev := devices.Registry["ws281x"]
	dev.Setup(pixels)

	colors := make([]uint32, pixels)

	for {
		for idx := range colors {
			colors[idx] = rand.Uint32()
		}
		dev.ShowColors(colors)
		//time.Sleep(time.Millisecond * 500)
	}
}
