package DTO

type RequestDTOType interface {
	DecrementRequestDTO | IncrementRequestDTO | DeleteRequestDTO | PutRequestDTO | GetRequestDTO | SaveToFileRequestDTO | RetrieveFromFileRequestDTO
}

type ResponseDTOType interface {
	DecrementResponseDTO | IncrementResponseDTO | DeleteResponseDTO | PutResponseDTO | GetResponseDTO | SaveToFileResponseDTO | RetrieveFromFileResponseDTO
}
