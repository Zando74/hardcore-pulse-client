package port

import "deathlog-tracker/domain/entity"

type GamePathRepository interface {
	Save(gamePath entity.GamePath) error
	Find() (*entity.GamePath, error)
}