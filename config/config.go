package config

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Width  float64
	Height float64
}

var Config Configuration

func Setup() {
	content, _ := ioutil.ReadFile("config/config.game.json")
	json.Unmarshal(content, &Config)
}
