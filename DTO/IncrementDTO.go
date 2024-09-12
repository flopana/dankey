package DTO

type IncrementRequestDTO struct {
	Database uint
	Key      string
}

type IncrementResponseDTO struct {
	ResponseDTO
	Value int `json:"value"`
}
