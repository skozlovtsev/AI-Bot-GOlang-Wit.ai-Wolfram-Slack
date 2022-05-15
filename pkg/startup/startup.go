package startup

import (
	"encoding/json"
	"io/ioutil"
)

var config *Config

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {}
	err = json.Unmarshal(file, &config)
	if err != nil {}
}