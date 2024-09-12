package main

import (
	"dankey/DTO"
	"dankey/HTTP"
	"dankey/Storage"
	"dankey/Storage/RAM"
)

func main() {
	var provider Storage.Provider = RAM.NewRamProvider()
	provider.Get(DTO.GetRequestDTO{
		Database: 1,
		Key:      "hi",
	})
	provider.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "hi",
		Value:    "hello",
	})

	var server = HTTP.NewServer(provider)
	server.Start()
}
