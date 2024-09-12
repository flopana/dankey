package RAM

import (
	"dankey/Storage/DTO"
	"dankey/Storage/RAM"
	"testing"
)

func TestIncrementIntRamProvider(t *testing.T) {
	var storage = RAM.NewRamProvider()
	var response = storage.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    1,
	})

	if response.Success == false {
		t.Errorf("Failed to inser value")
	}

	incrementResponse := storage.Increment(DTO.IncrementRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if incrementResponse.Success == false {
		t.Errorf("Failed to increment value")
	}

	if incrementResponse.Value != 2 {
		t.Errorf("Failed to increment value correctly")
	}
}

func TestIncrementStringRamProvider(t *testing.T) {
	var storage = RAM.NewRamProvider()
	var response = storage.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    "hello world",
	})

	if response.Success == false {
		t.Errorf("Failed to insert value")
	}

	incrementResponse := storage.Increment(DTO.IncrementRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if incrementResponse.Success == true {
		t.Errorf("Failed to return error for incrementing string")
	}
}
