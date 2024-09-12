package RAM

import "dankey/DTO"

func (provider *RamProvider) Get(request DTO.GetRequestDTO) DTO.GetResponseDTO {
	if _, ok := provider.storage[request.Database]; !ok {
		return DTO.GetResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Database does not exist",
			},
			Value: nil,
		}
	}

	if !provider.checkIfKeyExists(request.Database, request.Key) {
		return DTO.GetResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: "Key does not exist",
			},
			Value: nil,
		}
	}

	return DTO.GetResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Retrieved",
		},
		Value: provider.storage[request.Database][request.Key],
	}
}
