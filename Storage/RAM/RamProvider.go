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

func checkIfInt(value any) bool {
	switch value.(type) {
	case int:
		return true
	default:
		return false
	}
}
