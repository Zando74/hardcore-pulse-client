package factory

import (
	"deathlog-tracker/domain/entity"
	value_object "deathlog-tracker/domain/value-object"
)

type PositionFactory struct {}

func (p *PositionFactory) CreatePosition(mapID int, x float32, y float32) (*entity.Position, error) {

	xCoordinate, err := value_object.NewCoordinate(int(x * 100))
	if err != nil {
		return nil, err
	}

	yCoordinate, err := value_object.NewCoordinate(int(y * 100))
	if err != nil {
		return nil, err
	}
	
	mapId, err := value_object.NewMapID(mapID)
	if err != nil {
		return nil, err
	}

	return &entity.Position{
		MapID: *mapId,
		X: *xCoordinate,
		Y: *yCoordinate,
	}, nil
}