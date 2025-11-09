package value_object

type PlayerName string

type InvalidPlayerNameError struct {
	name string
}

func (e *InvalidPlayerNameError) Error() string {
	return "Invalid player name : " + e.name
}

func isValidPlayerName(name string) bool {
	runes := []rune(name)
	return len(runes) >= 2 && len(runes) <= 12
}

func NewPlayerName(value string) (*PlayerName, error) {
	if isValidPlayerName(value) {
		name := PlayerName(value)
		return &name, nil
	}
	return nil, &InvalidPlayerNameError{name: value}
}