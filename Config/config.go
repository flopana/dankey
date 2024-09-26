package Config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewConfig() (*Config, error) {
	return NewConfigWithPath("config.json")
}

func NewConfigWithPath(filePath string) (*Config, error) {
	config := Config{}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// unmashal the data
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
