package query

import (
	"deathlog-tracker/domain/entity"
	"deathlog-tracker/domain/port"
)

type GetGamePathQuery struct {}

type GetGamePathQueryHandler struct {
	GamePathRepository port.GamePathRepository
}

func (h *GetGamePathQueryHandler) Handle(query GetGamePathQuery) (*entity.GamePath, error) {

	gamePath, err := h.GamePathRepository.Find()
	if err != nil {
		return &entity.GamePath{ Path: ""}, nil
	}

	return gamePath, nil
}