package command

import (
	"deathlog-tracker/app/service"
	"deathlog-tracker/domain/entity"
	"deathlog-tracker/domain/factory"
	port "deathlog-tracker/domain/port"
	"log"
)

type ProcessAndSendBatchOfDeathCommand struct {
	PlayerBatch []factory.PlayerFactoryInput
}

type ProcessAndSendBatchOfDeathCommandHandler struct {
	PlayerHashRepository port.PlayerHashRepository
	PlayerSender port.PlayerSender
}


func (h *ProcessAndSendBatchOfDeathCommandHandler) Handle(command ProcessAndSendBatchOfDeathCommand) error {
	
	players := make([]entity.Player, 0)
	hashes := make([]string, 0)
	cpt := 0

	for _, player := range command.PlayerBatch {
		if cpt >= 1000 { // process only 1000 players at a time
			break
		}
		player, err := (*factory.PlayerFactory).CreatePlayer(nil, player)
		if err == nil {
			hash := service.Hash(*player)
			if !h.PlayerHashRepository.Exist(hash) {
				players = append(players, *player)
				hashes = append(hashes, hash)
				cpt++
			}
		}
	}
	if len(players) == 0 {
		log.Println("[DEATHLOG-TRACKER] No new player deaths to transfer")
		return nil
	}
	success := h.PlayerSender.SendBatch(players)
	if success {
		err := h.PlayerHashRepository.SaveAll(hashes)
		if err != nil {
			return err
		}
	}
	return nil
}