package RAM

import (
	"dankey/Storage/DTO"
	"dankey/Storage/RAM"
	"testing"
)

func TestPutAndGetRamProvider(t *testing.T) {
	var storage = RAM.NewRamProvider()
	var response = storage.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    "hello world",
	})

	if response.Success == false {
		t.Errorf("Failed to insert value")
	}

	getResponse := storage.Get(DTO.GetRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if getResponse.Success == false {
		t.Errorf("Failed to retrieve value")
	}

	if getResponse.Value != "hello world" {
		t.Errorf("Failed to retrieve correct value")
	}
}
