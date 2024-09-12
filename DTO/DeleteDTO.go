package DTO

type DeleteRequestDTO struct {
	Database uint
	Key      string
}

type DeleteResponseDTO struct {
	ResponseDTO
}
