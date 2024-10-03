package RAM

import (
	"dankey/Storage"
)

type RamProvider struct {
	storage map[uint]map[string]any
	Storage.Provider
}

func NewRamProvider() *RamProvider {
	return &RamProvider{
		storage: make(map[uint]map[string]any),
	}
}

func (provider *RamProvider) checkIfKeyExists(database uint, key string) bool {
	if _, ok := provider.storage[database][key]; ok {
		return true
	}
	return false
}

func (provider *RamProvider) checkIfDatabaseExists(datatabase uint) bool {
	if _, ok := provider.storage[datatabase]; ok {
		return true
	}
	return false
}

func checkIfInt(val any) (bool, any) {
	switch v := val.(type) {
	case int:
		return true, v
	case float64:
		if v == float64(int(v)) {
			return true, int(v) // Convert to int if no information is lost
		}
		return false, nil
	case float32:
		if v == float32(int(v)) {
			return true, int(v) // Convert to int if no information is lost
		}
		return false, nil
	default:
		return false, nil
	}
}
