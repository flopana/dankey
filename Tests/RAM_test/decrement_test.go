package RAM

import (
	"dankey/DTO"
	"dankey/Storage/RAM"
	"testing"
)

func TestDecrement(t *testing.T) {
	provider := RAM.NewRamProvider()
	provider.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    2,
	})

	response := provider.Decrement(DTO.DecrementRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if response.Success == false {
		t.Errorf("Failed to decrement value")
	}

	if response.Value != 1 {
		t.Errorf("Failed to decrement value correctly")
	}

}

func TestDecrementOfNotInt(t *testing.T) {
	provider := RAM.NewRamProvider()

	provider.Put(DTO.PutRequestDTO{
		Database: 0,
		Key:      "key",
		Value:    make([]byte, 10),
	})

	var response = provider.Decrement(DTO.DecrementRequestDTO{
		Database: 0,
		Key:      "key",
	})

	if response.Success == true {
		t.Errorf("Value decremented even though it is not an integer")
	}
}
