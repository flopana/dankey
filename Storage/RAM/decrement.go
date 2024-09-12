package RAM

import "dankey/DTO"

func (provider *RamProvider) Decrement(request DTO.DecrementRequestDTO) DTO.DecrementResponseDTO {
	if _, ok := provider.storage[request.Database]; !ok {
		provider.storage[request.Database] = make(map[string]any)
	}

	if !provider.checkIfKeyExists(request.Database, request.Key) {
		return DTO.DecrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key does not exist",
			},
			Value: 0,
		}
	}

	if checkIfInt(provider.storage[request.Database][request.Key]) == false {
		return DTO.DecrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key is not an integer",
			},
			Value: 0,
		}
	}

	provider.storage[request.Database][request.Key] = provider.storage[request.Database][request.Key].(int) - 1

	return DTO.DecrementResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Decremented",
		},
		Value: provider.storage[request.Database][request.Key].(int),
	}
}
