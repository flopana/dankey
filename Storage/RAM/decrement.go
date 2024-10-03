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

	value := provider.storage[request.Database][request.Key]
	ok, convertedValue := checkIfInt(value)

	if !ok {
		return DTO.DecrementResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key is not an integer",
			},
		}
	}

	provider.storage[request.Database][request.Key] = convertedValue.(int) - 1

	return DTO.DecrementResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Decremented",
		},
		Value: provider.storage[request.Database][request.Key].(int),
	}
}
