package value_object

import "slices"

type PlayerClass string

type InvalidPlayerClassError struct {
	class string
}

func (e *InvalidPlayerClassError) Error() string {
	return "Invalid player class : " + e.class
}

var ValidClasses = []string{
	"WARRIOR", 
	"PALADIN", 
	"HUNTER", 
	"MAGE", 
	"ROGUE", 
	"PRIEST", 
	"SHAMAN", 
	"WARLOCK", 
	"DRUID",
}

func NewPlayerClass(value string) (*PlayerClass, error) {
	if slices.Contains(ValidClasses, value) {
		class := PlayerClass(value)
		return &class, nil
	}
	return nil, &InvalidPlayerClassError{
		class: value,
	}
}