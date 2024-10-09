package RAM

func (provider *RamProvider) GetTotalKeys() uint64 {
	databases := provider.storage
	keyCount := uint64(0)
	for i := 0; i < len(databases); i++ {
		keyCount += uint64(len(databases[uint(i)]))
	}

	return keyCount
}

func (provider *RamProvider) GetTotalDatabases() uint64 {
	return uint64(len(provider.storage))
}
