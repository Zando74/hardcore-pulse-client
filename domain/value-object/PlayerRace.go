package value_object

import "slices"

type PlayerRace string

type InvalidPlayerRaceError struct {
	race string
}

func (e *InvalidPlayerRaceError) Error() string {
	return "Invalid player race : " + e.race
}

var ValidRaces = []string{
	"Human",
	"Orc",
	"Dwarf",
	"Night Elf",
	"Undead",
	"Tauren",
	"Gnome",
	"Troll",
}

func NewPlayerRace(value string) (*PlayerRace, error) {
	if slices.Contains(ValidRaces, value) {
		race := PlayerRace(value)
		return &race, nil
	}
	return nil, &InvalidPlayerRaceError{
		race: value,
	}
}