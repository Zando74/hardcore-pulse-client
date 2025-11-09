package command

import (
	"deathlog-tracker/domain/entity"
	"deathlog-tracker/domain/port"
)

type AskGamePathCommand struct {
	GetUserPath func() string
}

type AskGamePathCommandHandler struct {
	GamePathRepository port.GamePathRepository
}

func (h *AskGamePathCommandHandler) Handle(command AskGamePathCommand) error {
	path := command.GetUserPath()

	err := h.GamePathRepository.Save(entity.GamePath{ Path: path})
	if err != nil {
		return err
	}
	return nil
}