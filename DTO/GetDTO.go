package DTO

type GetRequestDTO struct {
	Database uint   `validate:"required"`
	Key      string `validate:"required"`
}

type GetResponseDTO struct {
	ResponseDTO
	Value any `json:"value"`
}
