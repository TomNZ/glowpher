package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tomnz/glowpher/internal/player"

	"github.com/k0kubun/pp"
	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/devices"
	"github.com/tomnz/glowpher/internal/dsl"
)

const pixels = 52

func main() {
	var cfg config.Config

	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(configFile, &cfg); err != nil {
		log.Fatalf("couldn't read dsl: %s", err)
	}

	pp.Print(cfg)
	fmt.Println()
	fmt.Println()

	dslPl, err := dsl.Compile(cfg)
	if err != nil {
		log.Fatalf("couldn't compile dsl: %s", err)
	}

	pp.Print(dslPl)
	fmt.Println()
	fmt.Println()

	cfgApi := dsl.Decompile(dslPl)

	pp.Print(cfgApi)
	fmt.Println()

	dev := devices.Registry["ws281x"]
	dev.Setup(pixels)

	for {
		player.Play(dslPl.Playlists["Christmas!"], dslPl.Scenes, dev, player.WithNumLeds(pixels))
	}
}
