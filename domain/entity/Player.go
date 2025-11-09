package entity

import value_object "deathlog-tracker/domain/value-object"

type Player struct {
	Name       value_object.PlayerName `json:"name"`
	Position   Position `json:"position"`
	Level      value_object.PlayerLevel `json:"level"`
	Class      value_object.PlayerClass `json:"class"`
	Race       value_object.PlayerRace `json:"race"`
	SourceID   string `json:"sourceID"`
	Guild      value_object.Guild `json:"guild"`
	Faction    value_object.Faction `json:"faction"`
	Realm      value_object.Realm `json:"realm"`
	LastWord   value_object.LastWord `json:"lastWord"`
	Timestamp int64 `json:"timestamp"`
}