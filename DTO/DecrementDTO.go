package DTO

type DecrementRequestDTO struct {
	Database uint
	Key      string
}

type DecrementResponseDTO struct {
	ResponseDTO
	Value int `json:"value"`
}
