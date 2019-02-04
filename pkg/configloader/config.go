package configloader

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

// Load reads the named JSON configuration file. This
// function fails if the file is not found or cannot be
// parsed. The loaded configuration is available in config.Vars
func Load(filePath string, configVar interface{}) {
	//filename is the path to the json config file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Cound not open config file", filePath, err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(configVar)
	if err != nil {
		log.Fatal("Could not parse config file", filePath, err)
	}
}
