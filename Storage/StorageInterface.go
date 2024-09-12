package Storage

import "dankey/DTO"

type Provider interface {
	Put(dto DTO.PutRequestDTO) DTO.PutResponseDTO
	Get(dto DTO.GetRequestDTO) DTO.GetResponseDTO
	Delete(dto DTO.DeleteRequestDTO) DTO.DeleteResponseDTO
	Increment(dto DTO.IncrementRequestDTO) DTO.IncrementResponseDTO
	Decrement(dto DTO.DecrementRequestDTO) DTO.DecrementResponseDTO
	SaveToFile(dto DTO.SaveToFileRequestDTO) DTO.SaveToFileResponseDTO
	RetrieveFromFile(dto DTO.RetrieveFromFileRequestDTO) DTO.RetrieveFromFileResponseDTO
}
