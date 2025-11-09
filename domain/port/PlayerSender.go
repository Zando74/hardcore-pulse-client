package port

import "deathlog-tracker/domain/entity"

type PlayerSender interface {
	SendBatch(players []entity.Player) bool
}