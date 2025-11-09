package factory

import (
	"deathlog-tracker/domain/entity"
	value_object "deathlog-tracker/domain/value-object"
	"strconv"
)

type PlayerFactoryInput struct {
	Name string
	MapID int
	X float32
	Y float32
	Level int
	Class string
	Race string
	Guild string
	Faction string
	Realm string
	SourceID int
	LastWord string
	Timestamp int64
}


type PlayerFactory struct {}

func (p *PlayerFactory) CreatePlayer(input PlayerFactoryInput) (*entity.Player, error) {
	name, err := value_object.NewPlayerName(input.Name)
	if err != nil {
		return nil, err
	}
	
	position, err := (*PositionFactory).CreatePosition(nil, input.MapID, input.X, input.Y)
	if err != nil {
		return nil, err
	}
	
	level, err := value_object.NewPlayerLevel(input.Level)
	if err != nil {
		return nil, err
	}
	class, err := value_object.NewPlayerClass(input.Class)
	if err != nil {
		return nil, err
	}
	race, err := value_object.NewPlayerRace(input.Race)
	if err != nil {
		return nil, err
	}
	guild, err := value_object.NewGuild(input.Guild)
	if err != nil {
		return nil, err
	}
	faction, err := value_object.NewFaction(input.Faction)
	if err != nil {
		return nil, err
	}
	realm, err := value_object.NewRealm(input.Realm)
	if err != nil {
		return nil, err
	}
	lastWord, err := value_object.NewLastWord(input.LastWord)
	if err != nil {
		return nil, err
	}

	player := entity.Player{
		Name: *name,
		Position: *position,
		Level: *level,
		Class: *class,
		Race: *race,
		Guild: *guild,
		Faction: *faction,
		Realm: *realm,
		SourceID: strconv.Itoa(input.SourceID),
		LastWord: *lastWord,
		Timestamp: input.Timestamp,
	}
	return &player, nil
}