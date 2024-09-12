package RAM_test

import (
	"dankey/Storage/DTO"
	"dankey/Storage/RAM"
	"testing"
)

func TestDelete(t *testing.T) {
	provider := RAM.NewRamProvider()
	provider.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    "hello world",
	})

	response := provider.Delete(DTO.DeleteRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if response.Success == false {
		t.Errorf("Failed to delete key")
	}

	if provider.Get(DTO.GetRequestDTO{
		Database: 0,
		Key:      "key",
	}).Success == true {
		t.Errorf("Retrieved the key after deletion")
	}
}
