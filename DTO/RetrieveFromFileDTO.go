package DTO

type RetrieveFromFileRequestDTO struct {
	FilePath string
}

type RetrieveFromFileResponseDTO struct {
	ResponseDTO
	Size              int64  `json:"size"`
	SizeHumanReadable string `json:"sizeHumanReadable"`
	FilePath          string `json:"filePath"`
}
