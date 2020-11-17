package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

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

	pl, err := dsl.Compile(cfg)
	if err != nil {
		log.Fatalf("couldn't compile dsl: %s", err)
	}

	pp.Print(pl)
	fmt.Println()
	fmt.Println()

	cfgApi := dsl.Decompile(pl)

	pp.Print(cfgApi)
	fmt.Println()

	dev := devices.Registry["ws281x"]
	dev.Setup(pixels)

	colors := make([]uint32, pixels)

	for {
		for idx := range colors {
			colors[idx] = rand.Uint32()
		}
		dev.ShowColors(colors)
		time.Sleep(time.Millisecond * 500)
	}
}
