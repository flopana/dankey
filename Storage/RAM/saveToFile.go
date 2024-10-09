package RAM

import (
	"dankey/DTO"
	"dankey/Util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func (provider *RamProvider) SaveToFile(dto DTO.SaveToFileRequestDTO) DTO.SaveToFileResponseDTO {
	provider.storageMutex.RLock()
	defer provider.storageMutex.RUnlock()
	file, err := os.OpenFile(dto.FilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return defaultSaveErrorResponse(err, &dto)
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return defaultSaveErrorResponse(err, &dto)
	}

	marshal, err := bson.Marshal(provider.storage)
	if err != nil {
		return defaultSaveErrorResponse(err, &dto)
	}

	_, err = file.Write(marshal)
	if err != nil {
		return defaultSaveErrorResponse(err, &dto)
	}

	fileStat, err := file.Stat()
	if err != nil {
		return defaultSaveErrorResponse(err, &dto)
	}

	res := DTO.SaveToFileResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Data saved to file",
		},
		Size:              fileStat.Size(),
		SizeHumanReadable: Util.ByteCountSI(fileStat.Size()),
		FilePath:          dto.FilePath,
	}

	log.Info().
		Int64("Size", res.Size).
		Str("SizeHumanReadable", res.SizeHumanReadable).
		Str("FilePath", res.FilePath).
		Msg("Data saved to file")

	return res
}

func defaultSaveErrorResponse(err error, dto *DTO.SaveToFileRequestDTO) DTO.SaveToFileResponseDTO {
	log.Error().Msg("Error saving data to file")
	log.Err(err).Msg("")
	return DTO.SaveToFileResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		},
		FilePath: dto.FilePath,
	}
}
