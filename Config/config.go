package Config

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func NewConfig() (*Config, error) {
	return NewConfigWithPath("config.json")
}

func NewConfigWithPath(filePath string) (*Config, error) {
	config := getDefaultConfig()
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn().Msg("Config file not found. Using default config.")
			log.Warn().Msg("Please create a config.json file in the root directory to customize the configuration.")
			log.Warn().Interface("Default config", config).Msg("")
			return config, nil
		}
		return nil, err
	}
	defer file.Close()

	// unmashal the data
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	// if credentials are the default ones, warn the user
	if config.Username == "admin" && config.Password == "admin" {
		log.Warn().Msg("Using default credentials. Please change them in the config file.")
	}

	return config, nil
}

func getDefaultConfig() *Config {
	return &Config{
		Username: "admin",
		Password: "admin",
		Port:     "6969",
	}
}
