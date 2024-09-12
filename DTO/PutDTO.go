package DTO

type PutRequestDTO struct {
	Database uint   `validate:"required"`
	Key      string `validate:"required"`
	Value    any    `validate:"required"`
}

type PutResponseDTO struct {
	ResponseDTO
}
