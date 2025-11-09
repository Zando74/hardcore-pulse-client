package port

import (
	"deathlog-tracker/domain/entity"
)

type FileReader interface {
	ExtractPlayerDeathLogData(filePath string) ([]entity.DeathRecord, error)
}