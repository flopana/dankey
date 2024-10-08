package RAM

import (
	"dankey/DTO"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func (provider *RamProvider) RetrieveFromFile(dto DTO.RetrieveFromFileRequestDTO) DTO.RetrieveFromFileResponseDTO {
	file, err := os.OpenFile(dto.FilePath, os.O_RDONLY, 0666)
	if err != nil {
		return defaultRetrieveErrorResponse(err, &dto)
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return defaultRetrieveErrorResponse(err, &dto)
	}

	// Read the file
	data := make([]byte, fileStat.Size())
	_, err = file.Read(data)

	// Unmarshal the data
	provider.storageMutex.Lock()
	defer provider.storageMutex.Unlock()
	err = bson.Unmarshal(data, &provider.storage)
	if err != nil {
		return defaultRetrieveErrorResponse(err, &dto)
	}

	res := DTO.RetrieveFromFileResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Data retrieved from file",
		},
		Size:              fileStat.Size(),
		SizeHumanReadable: byteCountSI(fileStat.Size()),
		FilePath:          dto.FilePath,
	}

	log.Info().
		Int64("Size", res.Size).
		Str("SizeHumanReadable", res.SizeHumanReadable).
		Str("FilePath", res.FilePath).
		Msg("Data retrieved from file")

	return res
}

func defaultRetrieveErrorResponse(err error, dto *DTO.RetrieveFromFileRequestDTO) DTO.RetrieveFromFileResponseDTO {
	return DTO.RetrieveFromFileResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		},
		FilePath: dto.FilePath,
	}
}
