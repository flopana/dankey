package RAM

import "dankey/DTO"

func (provider *RamProvider) Put(request DTO.PutRequestDTO) DTO.PutResponseDTO {
	provider.storageMutex.Lock()
	defer provider.storageMutex.Unlock()
	if !provider.checkIfDatabaseExists(request.Database) {
		provider.storage[request.Database] = make(map[string]any)
	}

	provider.storage[request.Database][request.Key] = request.Value

	return DTO.PutResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Inserted",
		},
	}
}
