package RAM

import "dankey/DTO"

func (provider *RamProvider) Increment(request DTO.IncrementRequestDTO) DTO.IncrementResponseDTO {
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

	if checkIfInt(provider.storage[request.Database][request.Key]) == false {
		return DTO.IncrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key is not an integer",
			},
			Value: 0,
		}
	}

	provider.storage[request.Database][request.Key] = provider.storage[request.Database][request.Key].(int) + 1

	return DTO.IncrementResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Incremented",
		},
		Value: provider.storage[request.Database][request.Key].(int),
	}
}
