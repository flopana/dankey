package RAM

import "dankey/DTO"

func (provider *RamProvider) Increment(request DTO.IncrementRequestDTO) DTO.IncrementResponseDTO {
	provider.storageMutex.Lock()
	defer provider.storageMutex.Unlock()
	if _, ok := provider.storage[request.Database]; !ok {
		provider.storage[request.Database] = make(map[string]any)
	}

	if !provider.checkIfKeyExists(request.Database, request.Key) {
		return DTO.IncrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key does not exist",
			},
			Value: 0,
		}
	}

	value := provider.storage[request.Database][request.Key]
	ok, convertedValue := checkIfInt(value)

	if !ok {
		return DTO.IncrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key is not an integer",
			},
		}
	}

	provider.storage[request.Database][request.Key] = convertedValue.(int) + 1

	return DTO.IncrementResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Incremented",
		},
		Value: provider.storage[request.Database][request.Key].(int),
	}
}
