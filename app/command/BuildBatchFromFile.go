package command

import (
	"deathlog-tracker/domain/factory"
	"deathlog-tracker/domain/port"
)

type BuildBatchFromFileCommand struct {
	FilePath string
}

type BuildBatchFromFileCommandHandler struct {
	FileReader port.FileReader
}

func (h *BuildBatchFromFileCommandHandler) Handle(command BuildBatchFromFileCommand) ([]factory.PlayerFactoryInput, error) {
	
	deathRecord, err := h.FileReader.ExtractPlayerDeathLogData(command.FilePath)
	if err != nil {
		return nil, err
	}

	playerBatch := make([]factory.PlayerFactoryInput, len(deathRecord))

	for i, deathRecord := range deathRecord {
		playerBatch[i] = factory.PlayerFactoryInput{
			Name: deathRecord.Name,
			MapID: deathRecord.GetMapID(),
			X: deathRecord.GetXPos(),
			Y: deathRecord.GetYPos(),
			Level: deathRecord.Level,
			Class: deathRecord.GetClassName(),
			Race: deathRecord.GetRaceName(),
			SourceID: deathRecord.SourceID,
			Guild: deathRecord.Guild,
			Faction: deathRecord.GetFactionName(),
			Realm: deathRecord.Realm,
			LastWord: deathRecord.LastWords,
			Timestamp: deathRecord.Date,
		}
	}
	
	return playerBatch, nil
}