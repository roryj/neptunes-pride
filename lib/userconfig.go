package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultConfigPath = "./config.json"

type UserConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseConfigFile() UserConfig {
	return ParseConfigFileWithPath(defaultConfigPath)
}

func ParseConfigFileWithPath(filePath string) UserConfig {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var config UserConfig

	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&config); err != nil {
		panic(fmt.Sprintf("Error parsing config file. %v", err.Error()))
	}

	return config
}
