package main

import (
	"dankey/Config"
	"dankey/HTTP"
	"dankey/Storage"
	"dankey/Storage/RAM"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	configureLogger()
	var config, err = Config.NewConfig()
	if err != nil {
		log.Fatal().Msg("Error reading config file. Aborting.")
	}
	log.Info().Msg("Config loaded successfully")

	var provider Storage.Provider = RAM.NewRamProvider()
	var server = HTTP.NewServer(provider, config)
	server.Start()
}

func configureLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
}
