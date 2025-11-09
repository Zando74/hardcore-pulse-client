package entity

import value_object "deathlog-tracker/domain/value-object"

type Position struct {
	MapID value_object.MapID `json:"mapID"`
	X value_object.Coordinate `json:"x"`
	Y value_object.Coordinate `json:"y"`
}