package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/playlist"
)

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
}
