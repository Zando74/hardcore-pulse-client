package value_object

import "fmt"

type Coordinate int

type InvalidCoordinateError struct {
	coord int
}

func (e *InvalidCoordinateError) Error() string {
	return "Invalid coordinate"  + fmt.Sprint(e.coord)
}

func NewCoordinate(value int) (*Coordinate, error) {
	if value >= 0 && value <= 100 {
		coord := Coordinate(value)
		return &coord, nil
	}
	return nil, &InvalidCoordinateError{
		coord: value,
	}
}