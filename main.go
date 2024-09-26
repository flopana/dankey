package main

import (
	"dankey/Config"
	"dankey/HTTP"
	"dankey/Storage"
	"dankey/Storage/RAM"
	"fmt"
)

func main() {
	var config, err = Config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	var provider Storage.Provider = RAM.NewRamProvider()

	var server = HTTP.NewServer(provider, config)
	server.Start()
}
