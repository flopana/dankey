package RAM

import "dankey/DTO"

func (provider *RamProvider) Delete(request DTO.DeleteRequestDTO) DTO.DeleteResponseDTO {
	if ok := provider.checkIfKeyExists(request.Database, request.Key); ok {
		delete(provider.storage[request.Database], request.Key)
		return DTO.DeleteResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: true,
				Message: "Key deleted",
			},
		}
	}
	return DTO.DeleteResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: false,
			Message: "Key does not exist",
		},
	}
}
