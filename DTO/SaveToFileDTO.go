package DTO

type SaveToFileRequestDTO struct {
	FilePath string
}

type SaveToFileResponseDTO struct {
	ResponseDTO
	Size              int64  `json:"size"`
	SizeHumanReadable string `json:"sizeHumanReadable"`
	FilePath          string `json:"filePath"`
}
