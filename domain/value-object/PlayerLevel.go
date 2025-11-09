package value_object

import "fmt"

type PlayerLevel int

type InvalidPlayerLevelError struct {
	level int
}

func (e *InvalidPlayerLevelError) Error() string {
	return "Invalid player level : " + fmt.Sprint(e.level)
}

func NewPlayerLevel(value int) (*PlayerLevel, error) {
	if value >= 1 && value <= 60 {
		level := PlayerLevel(value)
		return &level, nil
	}
	return nil, &InvalidPlayerLevelError{
		level: value,
	}
}