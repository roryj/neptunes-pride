package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultConfigPath = "./config.json"

type GameConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GameId   string `json:"game_id"`
}

func ParseConfigFile() GameConfig {
	return ParseConfigFileWithPath(defaultConfigPath)
}

func ParseConfigFileWithPath(filePath string) GameConfig {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var config GameConfig

	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&config); err != nil {
		panic(fmt.Sprintf("Error parsing config file. %v", err.Error()))
	}

	return config
}
