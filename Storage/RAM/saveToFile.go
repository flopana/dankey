package RAM

import (
	"dankey/DTO"
	"encoding/gob"
	"fmt"
	"os"
)

func (provider *RamProvider) SaveToFile(dto DTO.SaveToFileRequestDTO) DTO.SaveToFileResponseDTO {
	file, err := os.OpenFile(dto.FilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return DTO.SaveToFileResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: err.Error(),
			},
			FilePath: dto.FilePath,
		}
	}
	defer file.Close()

	var encoder = gob.NewEncoder(file)
	err = encoder.Encode(provider.storage)
	if err != nil {
		return DTO.SaveToFileResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: err.Error(),
			},
			FilePath: dto.FilePath,
		}
	}

	fileStat, err := file.Stat()
	if err != nil {
		return DTO.SaveToFileResponseDTO{
			ResponseDTO: DTO.ResponseDTO{
				Success: false,
				Message: err.Error(),
			},
			FilePath: dto.FilePath,
		}
	}

	return DTO.SaveToFileResponseDTO{
		ResponseDTO: DTO.ResponseDTO{
			Success: true,
			Message: "Data saved to file",
		},
		Size:              fileStat.Size(),
		SizeHumanReadable: byteCountSI(fileStat.Size()),
		FilePath:          dto.FilePath,
	}
}

/*
*
https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
*/
func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
