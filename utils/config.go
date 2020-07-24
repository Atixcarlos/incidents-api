package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the app config data
var Config ConfigData

// ConfigData defines the config data
type ConfigData struct {
	LogDirectory       string `json:"logDirectory"`
	NetworkAddr        string `json:"networkAddr"`
	DBConnectionString string `json:"dbConnectionString"`
}

// LoadConfig loads the data from a config file
func LoadConfig(file string) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config)
}
